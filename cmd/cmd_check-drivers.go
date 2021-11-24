package cmd

import (
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common"
	"github.com/spf13/cobra"
)

type DriverCheckOptions struct {
	*common.CommonOptions
	Output string
}

func NewCmdDriverCheck(commonOpts *common.CommonOptions) *cobra.Command {
	options := &InfoOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "check-drivers",
		Aliases: []string{"cd", "checkdrivers"},
		Short:   "Check for Driver Updates (BETA)",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}

	return cmd
}

func (o *DriverCheckOptions) Run() error {

	return nil
}
