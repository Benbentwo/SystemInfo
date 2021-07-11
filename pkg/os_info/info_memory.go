package os_info

import (
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/jaypipes/ghw"
	"github.com/jedib0t/go-pretty/v6/table"
)

func getMem() *ghw.MemoryInfo {
	memory, err := ghw.Memory()
	if err != nil {
		log.Logger().Errorf("Error getting memory info: %v", err)
	}
	return memory
}

func (sysInfo *SystemInformation) outputMemoryToTable() {
	sysInfo.Table.AppendHeader(table.Row{HEADER, "Memory:", HEADER})

	for _, module := range sysInfo.Memory.Modules {
		sysInfo.Table.AppendRow(
			table.Row{SPACE, module.Vendor, sysInfo.Memory.TotalPhysicalBytes},
		)
	}
	sysInfo.Table.AppendSeparator()
}
