package os_info

import (
	"github.com/jaypipes/ghw"
	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
	"os"
)

const (
	HEADER = "**********"
	SPACE  = ""
)

var SystemInfoWriter table.Writer = table.NewWriter()

type SystemInformation struct {
	Cpu      *ghw.CPUInfo    `json:"cpu"`
	Memory   *ghw.MemoryInfo `json:"memory"`
	Graphics *ghw.GPUInfo    `json:"graphics"`
	Disks    *ghw.BlockInfo  `json:"disks"`
}

func NewSysInfo(out terminal.FileWriter) *SystemInformation {
	sysInfo := &SystemInformation{}
	SystemInfoWriter.Style()
	SystemInfoWriter.SetOutputMirror(os.Stdout)

	return sysInfo
}

func (sysInfo *SystemInformation) GetAllInformation() {
	sysInfo.Cpu = getCpu()
	sysInfo.Memory = getMem()
	sysInfo.Graphics = getGpu()
	sysInfo.Disks = getDisk()
}

func (sysInfo *SystemInformation) RenderInformation() {
	sysInfo.outputCPUToTable()
	sysInfo.outputGPUToTable()
	sysInfo.outputMemoryToTable()
	sysInfo.outputDiskToTable()
	SystemInfoWriter.Render()
}
