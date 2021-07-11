package cmd

import (
	"code.cloudfoundry.org/bytefmt"
	"errors"
	"fmt"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/jaypipes/ghw"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
	"os"
)

type InfoOptions struct {
	*common.CommonOptions
	Output string
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

	cmd.Flags().StringVarP(&options.Output, "output", "o", "table", "How would you like to output the information?")

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
	fmt.Print(sysInfo.Cpu.YAMLString())
	//for _, processor := range sysInfo.Cpu.Processors {
	//	sysInfo.Table.AppendRow(
	//		table.Row{SPACE, processor.Vendor, ""},
	//	)
	//}
	sysInfo.Table.AppendSeparator()
}

func (sysInfo *SystemInformation) outputDiskToTable() {
	sysInfo.Table.AppendHeader(table.Row{HEADER, "DISKs:", HEADER})

	sysInfo.Table.AppendRow(
		table.Row{"TOTAL: ", SPACE, SPACE, SPACE, SPACE, bytefmt.ByteSize(sysInfo.Disks.TotalPhysicalBytes)},
	)
	for _, disk := range sysInfo.Disks.Disks {
		sysInfo.Table.AppendRow(
			table.Row{SPACE, disk.Name, disk.Vendor, bytefmt.ByteSize(disk.SizeBytes)},
		)
		for _, partition := range disk.Partitions {
			sysInfo.Table.AppendRow(
				table.Row{SPACE, SPACE, partition.Name, partition.Label, partition.Type, bytefmt.ByteSize(partition.SizeBytes)},
			)
		}
	}
	sysInfo.Table.AppendSeparator()
}

func (sysInfo *SystemInformation) RenderInformation() {
	sysInfo.outputCPUToTable()
	//sysInfo.outputMemoryToTable()
	sysInfo.outputDiskToTable()
	sysInfo.Table.Render()
}

func (o *InfoOptions) validateFlags() error {
	if o.Output != "table" && o.Output != "json" && o.Output != "yaml" {
		return errors.New("output arg must be one of `table` `json` or `yaml`")
	}

	// If we want output json or yaml, only print error messages on failure.
	if o.Output == "json" || o.Output == "yaml" {
		err := log.SetLevel("FATAL")
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *InfoOptions) Run() error {
	err := o.validateFlags()
	if err != nil {
		return err
	}

	sysInfo := NewSysInfo(o.Out)
	sysInfo.GetAllInformation()
	o.OutputInformation(sysInfo)
	return nil
}

func (o *InfoOptions) OutputInformation(information *SystemInformation) {
	if o.Output == "table" {
		information.RenderInformation()
	} else if o.Output == "json" {
		// Indent arg here formats the json vs minifies it. setting to false as it should be piped into `jq` or other
		//   formats for human readable.
		fmt.Print(information.Disks.JSONString(false))
	} else if o.Output == "yaml" {
		fmt.Print(information.Disks.YAMLString())
	}
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
