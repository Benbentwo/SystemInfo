package os_info

import (
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/jaypipes/ghw"
)

func GetMem() *ghw.MemoryInfo {
	memory, err := ghw.Memory()
	if err != nil {
		log.Logger().Errorf("Error getting memory info: %v", err)
	}
	return memory
}
