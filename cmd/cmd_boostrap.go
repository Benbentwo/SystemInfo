package cmd

import (
	"errors"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/bootstrap"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/utils"
	"github.com/spf13/cobra"
)

type BootstrapOptions struct {
	*common.CommonOptions
	Output  string
	Profile string
	Files   []string
	File    string
}

func NewCmdBootstrap(commonOpts *common.CommonOptions) *cobra.Command {
	options := &BootstrapOptions{
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
	cmd.Flags().StringVarP(&options.Profile, "profile", "p", "", "Specific Profile to install, defaults to all of them.")
	cmd.Flags().StringSliceVarP(&options.Files, "files", "f", []string{}, "List of files to combine before installing. repeat option or comma separated")

	return cmd
}

func (o *BootstrapOptions) validateFlags() error {
	if len(o.Files) == 0 && o.File == "" {
		return errors.New("You must input one file, use `-f`")
	}

	return nil
}

func (options *BootstrapOptions) Run() error {
	log.Logger().Debugf("Runnign with Files: %s", utils.ColorAnswer(options.Files))
	err := options.validateFlags()
	if err != nil {
		log.Logger().Errorf("%s", err)
		return errors.New("Argument Error")
	}

	for _, file := range options.Files {
		config, err := bootstrap.LoadBootstrapConfig(file)
		if err != nil {
			return err
		}

		log.Logger().Infof("Config Loaded: %v", config)
	}

	return nil
}
