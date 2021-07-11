package cmd

import (
	"errors"
	"fmt"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/os_info"
	"github.com/spf13/cobra"
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

	sysInfo := os_info.NewSysInfo(o.Out)
	sysInfo.GetAllInformation()
	o.OutputInformation(sysInfo)
	return nil
}

func (o *InfoOptions) OutputInformation(information *os_info.SystemInformation) {
	if o.Output == "table" {
		information.RenderInformation()
	} else if o.Output == "json" {
		// Indent arg here formats the json vs minifies it. setting to false as it should be piped into `jq` or other
		//   formats for human readable.

		// TODO - outputting each of these to json will create an invalid json...
		// 		  need to nest under system_info or as an array
		fmt.Print(information.Cpu.JSONString(false))
		fmt.Print(information.Memory.JSONString(false))
		fmt.Print(information.Graphics.JSONString(false))
		fmt.Print(information.Disks.JSONString(false))

	} else if o.Output == "yaml" {
		fmt.Print(information.Cpu.YAMLString())
		fmt.Print(information.Memory.YAMLString())
		fmt.Print(information.Graphics.YAMLString())
		fmt.Print(information.Disks.YAMLString())
	}
}
