package cmd

import (
	"errors"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/bootstrap"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/utils"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
	"strings"
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

	config := []bootstrap.Profile{}
	for _, file := range options.Files {
		newConfig, err := bootstrap.LoadBootstrapConfig(file)
		if err != nil {
			return err
		}

		config = append(config, newConfig...)
	}
	log.Logger().Infof("Config(s) Loaded: %v", config)
	// do some deep merging

	if options.BatchMode && options.Profile == "" && len(config) > 1 {
		//log.Logger().Errorf("Detected multiple profiles")
		return errors.New("Batch mode must specify a profile to install from `-p`, as multiple profiles were detected")
	}
	if !options.BatchMode {

		choices := make([]string, 0)
		for _, prof := range config {
			choices = append(choices, prof.ProfileName)
		}
		pick, err := util.Pick("Pick a profile to install:", choices, config[0].ProfileName)
		if err != nil {
			return err
		}
		options.Profile = pick
	}

	chosenProfile := bootstrap.Profile{}
	for _, conf := range config {
		if conf.ProfileName == options.Profile {
			chosenProfile = conf
			break
		}
	}

	for _, installer := range chosenProfile.Installers {
		packageManager, err := installer.GetInstaller()
		if err != nil {
			return err
		}
		installed, err := packageManager.IsInstalled()
		if err != nil {
			return err
		}
		if !installed {
			packageManager.InstallManager()
		}

		log.BeginSubCommandLogging(strings.ToUpper(installer.Name))
		packageManager.InstallPackages(installer.Packages)
		log.EndSubCommandLogging()
	}
	return nil
}
