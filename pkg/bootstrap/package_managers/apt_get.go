package package_managers

import (
	"github.com/Benbentwo/SystemInfo/pkg/common/log"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type Apt struct {
	//Cask	bool
}

func (a Apt) IsInstalled() (bool, error) {
	log.BeginSubCommandLogging("APT_GET_VERSION")
	cmd := exec.Command("apt-get", "--version")
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

func (a Apt) InstallManager() (string, error) {
	//log.Logger().WithFields(logrus.Fields{"pkg": pkg}).Error(err)
	//log.Logger().WithFields(logrus.Fields{"cmd": name, "args": args}).Debugf()
	panic("implement me")
	return "", nil
}

func (a Apt) InstallPackage(p string, options ...string) (bool, error) {
	log.Logger().WithFields(logrus.Fields{"cmd": "apt-get", "args": options, "pkg": p}).Infof("Installing: %s", p)

	//log.Logger().WithFields(logrus.Fields{"cmd": "brew", "args": options}).Debugf("Running ")
	log.BeginSubCommandLogging("APT_GET")
	stdout, err := Execute("apt-get", "install", "-y", strings.Join(options, " "), p)
	if err != nil {
		return false, err
	}
	log.Logger().WithFields(logrus.Fields{"args": options, "pkg": p}).Infof(stdout)
	return true, nil
}

func (a Apt) InstallPackages(p []string, options ...string) (map[string]PkgInstalledResult, error) {
	mapPkgs := make(map[string]PkgInstalledResult, 0)
	for _, pkg := range p {
		installed, err := a.InstallPackage(pkg, options...)
		if err != nil {
			mapPkgs[pkg] = PkgInstalledResult{Installed: false, Err: err}
			break
		}
		mapPkgs[pkg] = PkgInstalledResult{Installed: installed, Err: nil}
	}
	return mapPkgs, nil
}

//
//func (a Apt) InstallPackage() error {
//	return nil
//}
//
//func (a Apt) IsInstalled() error {

//}

//    sudo xcode-select --install
//    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
