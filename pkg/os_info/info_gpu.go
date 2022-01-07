package os_info

import (
	"github.com/Benbentwo/SystemInfo/pkg/common/log"
	"github.com/jaypipes/ghw"
	"github.com/jedib0t/go-pretty/v6/table"
	"runtime"
)

func getGpu() *ghw.GPUInfo {
	if runtime.GOOS == "darwin" {
		//gpuInfo, err := darwin.GetGpuDarwin()
		//if err != nil {
		//	log.Logger().Errorf("Failed to get GPU Information (OS: Darwin): %s", err)
		//}
		//return gpuInfo
	}
	gpu, err := ghw.GPU()
	if err != nil {
		log.Logger().Errorf("Error getting gpu info: %v", err)
	}
	return gpu
}

func (sysInfo *SystemInformation) outputGPUToTable() {
	if sysInfo.Graphics != nil {
		SystemInfoWriter.AppendRow(table.Row{HEADER, "Graphics:", HEADER})
		for _, graphicsCard := range sysInfo.Graphics.GraphicsCards {
			SystemInfoWriter.AppendRow(
				table.Row{SPACE, graphicsCard.DeviceInfo.Revision, SPACE, SPACE, SPACE, graphicsCard.DeviceInfo.Address},
			)
		}
		SystemInfoWriter.AppendSeparator()
	}
}
