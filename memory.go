package istockphoto

import (
	"github.com/spf13/viper"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	//MemoryYaml is filename of the memory digest file
	MemoryYaml = "_memory.yaml"
	// MemoryPlaceholder is the value placeholder in the `_memory.yaml` file
	MemoryPlaceholder = "_"
)

type memory struct {
	// Placeholder is the value placeholder in the `_memory.yaml` file
	Placeholder string
	// PathMemory is the relative path to the `_memory.yaml` file
	PathMemory string
	// Viper object holds the data container for the `download progress object`
	Viper *viper.Viper
}

func NewMemory(dirMemory string) *memory {
	m := &memory{
		Placeholder: MemoryPlaceholder,
		PathMemory:  filepath.Join(dirMemory, MemoryYaml),
	}
	m.Init()
	return m
}

func parseIstockID(s string) string {
	if strings.HasPrefix(s, "https://") {
		urlParse, _ := url.Parse(s)
		return urlParse.Query()["m"][0]
	} else if filepath.Ext(s) == ".jpg" {
		return strings.Split(s, "_")[1]
	} else {
		return s
	}
}

func (m *memory) Init() {
	if err := os.MkdirAll(filepath.Dir(m.PathMemory), os.ModePerm); err != nil {
		log.Println("Failed to create memory path: ", err)
		return
	}
	m.Viper = viper.New()
	m.Viper.SetConfigFile(m.PathMemory)
	m.loadMemory()
}

func (m *memory) DumpMemory() {
	if err := m.Viper.WriteConfigAs(m.PathMemory); err != nil {
		log.Println("Failed to dump memory: ", err)
	}
}

func (m *memory) SetMemory(k string) {
	m.Viper.Set(parseIstockID(k), m.Placeholder)
}

func (m *memory) GetMemory(k string) string {
	return m.Viper.GetString(parseIstockID(k))
}

func (m *memory) loadMemory() {
	dirMemory := filepath.Dir(m.PathMemory)
	files, _ := os.ReadDir(dirMemory)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".jpg" {
			m.SetMemory(file.Name())
		}
	}
}
