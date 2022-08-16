package istockphoto

import (
	"log"
	"os"
	"path"
)

const (
	//MemoryYaml 记忆体摘要文件 filename
	MemoryYaml = "_memory.yaml"
)

type Memory struct {
	// IstockDatabase `flag下载包` 路径
	IstockDatabase string

	// pathMemory 每个 `flag下载包` 的 `_memory.yaml` 文件路径
	pathMemory string

	// memory 存放 `下载进度对象` 的数据容器
	memory []map[string]interface{}
}

// InitMemory 初始化记忆体对象
func InitMemory(flag string) *Memory {
	memory := &Memory{
		IstockDatabase: flag,
	}

	// 补全记忆体摘要文件的绝对路径
	memory.pathMemory = path.Join(memory.IstockDatabase, MemoryYaml)
	err := os.MkdirAll(path.Dir(memory.pathMemory), 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatalln("MkdirAll err: ", err)
	}

	// 读取已有记忆
	memory.memory = memory.loadMemory()

	return memory
}

// DumpMemory 合并记忆体集合对象并将合并后的内容写回 `_memory.yaml`
func (m *Memory) DumpMemory(container []map[string]interface{}) {
}

// loadMemory 记忆体初始化时读取进行时的记忆体集合对象
func (m *Memory) loadMemory() []map[string]interface{} {
	return nil
}
