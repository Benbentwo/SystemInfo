package cmd

import (
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common"
	"github.com/spf13/cobra"
)

type BootstrapOptions struct {
	*common.CommonOptions
	Output string
}

func NewCmdBootstrap(commonOpts *common.CommonOptions) *cobra.Command {
	options := &InfoOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "bootstrap",
		Aliases: []string{"bs", "setup"},
		Short:   "Bootstrap this machine (BETA)",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}

	return cmd
}

func (o *BootstrapOptions) Run() error {

	return nil
}
