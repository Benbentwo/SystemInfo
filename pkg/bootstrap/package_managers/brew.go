package package_managers

import (
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type Brew struct {
	//Cask	bool
}

func (b Brew) IsInstalled() (bool, error) {
	log.BeginSubCommandLogging("BREW_VERSION")
	cmd := exec.Command("brew", "--version")
	stdout, err := cmd.Output()

	if err != nil {
		log.Logger().Errorf(err.Error())
		return false, err
	}

	// Print the output
	log.Logger().Infof(string(stdout))
	log.EndSubCommandLogging()
	return true, nil
}

func (b Brew) InstallManager() (string, error) {
	//log.Logger().WithFields(logrus.Fields{"pkg": pkg}).Error(err)
	//log.Logger().WithFields(logrus.Fields{"cmd": name, "args": args}).Debugf()
	panic("implement me")
	return "", nil
}

func (b Brew) InstallPackage(p string, options ...string) (bool, error) {
	log.Logger().WithFields(logrus.Fields{"cmd": "brew", "args": options, "pkg": p}).Infof("Installing: %s", p)

	//log.Logger().WithFields(logrus.Fields{"cmd": "brew", "args": options}).Debugf("Running ")
	log.BeginSubCommandLogging("BREW")
	stdout, err := Execute("brew", "install", strings.Join(options, " "), p)
	if err != nil {
		return false, err
	}
	log.Logger().WithFields(logrus.Fields{"args": options, "pkg": p}).Infof(stdout)
	return true, nil
}

func (b Brew) InstallPackages(p []string, options ...string) (map[string]PkgInstalledResult, error) {
	mapPkgs := make(map[string]PkgInstalledResult, 0)
	for _, pkg := range p {
		installed, err := b.InstallPackage(pkg, options...)
		if err != nil {
			mapPkgs[pkg] = PkgInstalledResult{Installed: false, Err: err}
			break
		}
		mapPkgs[pkg] = PkgInstalledResult{Installed: installed, Err: nil}
	}
	return mapPkgs, nil
}

//
//func (b Brew) InstallPackage() error {
//	return nil
//}
//
//func (b Brew) IsInstalled() error {

//}

//    sudo xcode-select --install
//    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
