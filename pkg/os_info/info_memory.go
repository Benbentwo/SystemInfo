package os_info

import (
	"code.cloudfoundry.org/bytefmt"
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
	SystemInfoWriter.AppendRow(table.Row{HEADER, "Memory:", HEADER})
	SystemInfoWriter.AppendRow(table.Row{"TOTAL Memory: ", SPACE, SPACE, SPACE, SPACE, bytefmt.ByteSize(uint64(sysInfo.Memory.TotalPhysicalBytes))})
	SystemInfoWriter.AppendRow(table.Row{"TOTAL Usable Memory: ", SPACE, SPACE, SPACE, SPACE, bytefmt.ByteSize(uint64(sysInfo.Memory.TotalUsableBytes))})

	for _, module := range sysInfo.Memory.Modules {
		SystemInfoWriter.AppendRow(
			table.Row{SPACE, module.Vendor, sysInfo.Memory.TotalPhysicalBytes},
		)
	}
	SystemInfoWriter.AppendSeparator()
}
