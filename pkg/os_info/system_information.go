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

type SystemInformation struct {
	Table    table.Writer
	Cpu      *ghw.CPUInfo
	Memory   *ghw.MemoryInfo
	Graphics *ghw.GPUInfo
	Disks    *ghw.BlockInfo
}

func NewSysInfo(out terminal.FileWriter) *SystemInformation {
	sysInfo := &SystemInformation{}
	sysInfo.Table = table.NewWriter()
	sysInfo.Table.Style()

	sysInfo.Table.SetOutputMirror(os.Stdout)
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
	sysInfo.outputMemoryToTable()
	sysInfo.outputDiskToTable()
	sysInfo.Table.Render()
}
