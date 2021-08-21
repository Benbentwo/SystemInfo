package os_info

import (
	"fmt"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/jaypipes/ghw"
	"github.com/jedib0t/go-pretty/v6/table"
)

func getGpu() *ghw.GPUInfo {
	gpu, err := ghw.GPU()
	if err != nil {
		log.Logger().Errorf("Error getting gpu info: %v", err)
	}
	return gpu
}

func (sysInfo *SystemInformation) outputGPUToTable() {
	SystemInfoWriter.AppendRow(table.Row{HEADER, "Graphics:", HEADER})
	fmt.Print(sysInfo.Graphics.YAMLString())
	for _, graphicsCard := range sysInfo.Graphics.GraphicsCards {
		SystemInfoWriter.AppendRow(
			table.Row{SPACE, graphicsCard.DeviceInfo.Revision, ""},
		)
	}
	SystemInfoWriter.AppendSeparator()
}
