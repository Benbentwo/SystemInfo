package cmd

import (
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/jaypipes/ghw"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
)

type InfoOptions struct {
	*common.CommonOptions
}

func NewCmdInfo(commonOpts *common.CommonOptions) *cobra.Command {
	options := &InfoOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:   "info",
		Short: "Print System Information",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	return cmd
}

type SystemInformation struct {
	Table    table.Writer
	Cpu      *ghw.CPUInfo
	Memory   *ghw.MemoryInfo
	Graphics *ghw.GPUInfo
	Disks    *ghw.BlockInfo
}

const (
	HEADER = "**********"
	SPACE  = ""
)

func NewSysInfo(out terminal.FileWriter) *SystemInformation {
	sysInfo := &SystemInformation{}
	sysInfo.Table = table.NewWriter()
	sysInfo.Table.SetOutputMirror(out)
	return sysInfo
}

func (sysInfo *SystemInformation) GetAllInformation() {
	sysInfo.Cpu = getCpu()
	sysInfo.Memory = getMem()
	sysInfo.Graphics = getGpu()
	sysInfo.Disks = getDisk()
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

func (sysInfo *SystemInformation) outputCPUToTable() {
	sysInfo.Table.AppendHeader(table.Row{HEADER, "CPU:", HEADER})

	for _, processor := range sysInfo.Cpu.Processors {
		sysInfo.Table.AppendRow(
			table.Row{SPACE, processor.Vendor, ""},
		)
	}
	sysInfo.Table.AppendSeparator()
}

func (sysInfo *SystemInformation) RenderInformation() {
	sysInfo.Table.Render()
}
func (o *InfoOptions) Run() error {
	sysInfo := NewSysInfo(o.Out)
	sysInfo.GetAllInformation()
	sysInfo.RenderInformation()
	return nil
}

func getMem() *ghw.MemoryInfo {
	memory, err := ghw.Memory()
	if err != nil {
		log.Logger().Errorf("Error getting memory info: %v", err)
	}
	return memory
}

func getCpu() *ghw.CPUInfo {
	cpu, err := ghw.CPU()
	if err != nil {
		log.Logger().Errorf("Error getting cpu info: %v", err)
	}
	return cpu
}
func getGpu() *ghw.GPUInfo {
	gpu, err := ghw.GPU()
	if err != nil {
		log.Logger().Errorf("Error getting gpu info: %v", err)
	}
	return gpu
}

func getDisk() *ghw.BlockInfo {
	disk, err := ghw.Block()
	if err != nil {
		log.Logger().Errorf("Error getting block info: %v", err)
	}
	return disk
}
