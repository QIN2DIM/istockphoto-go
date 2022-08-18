package downloader

import (
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	//memoryYaml is filename of the memory digest file
	memoryYaml = "_memory.yaml"
	// memoryPlaceholder is the value placeholder in the `_memory.yaml` file
	memoryPlaceholder = "_"
	// memorySuffix cached image format
	memorySuffix = ".jpg"
)

type memory struct {
	// Placeholder is the value placeholder in the `_memory.yaml` file
	Placeholder string
	// PathMemory is the relative path to the `_memory.yaml` file
	PathMemory string
	// ext cached image format
	ext string
	// container is cached images information
	container map[string]string
}

func newMemory(dirMemory string) *memory {
	m := &memory{
		PathMemory: filepath.Join(dirMemory, memoryYaml),
	}
	m.init()
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

func (m *memory) init() {
	m.Placeholder = memoryPlaceholder
	m.ext = memorySuffix
	m.container = make(map[string]string)
	if err := os.MkdirAll(filepath.Dir(m.PathMemory), os.ModePerm); err != nil {
		log.Println("Failed to create memory path: ", err)
		return
	}
	m.loadMemory()
}

// loadMemory read cached filenames and tokenize the data
func (m *memory) loadMemory() {
	dirMemory := filepath.Dir(m.PathMemory)
	files, _ := os.ReadDir(dirMemory)

	for _, file := range files {
		if filepath.Ext(file.Name()) == m.ext {
			m.setMemory(file.Name())
		}
	}
}

// GetMemory query memory
func (m *memory) GetMemory(k string) string {
	return m.container[parseIstockID(k)]
}

func (m *memory) setMemory(k string) {
	m.container[parseIstockID(k)] = m.Placeholder
}
