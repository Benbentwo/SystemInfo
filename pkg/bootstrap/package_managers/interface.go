package package_managers

import (
	"github.com/pkg/errors"
	"strings"
)

type PkgInstalledResult struct {
	Installed bool
	Err       error
}
type PackageManager interface {
	IsInstalled() (bool, error)
	InstallManager() (string, error)
	InstallPackage(p string, options ...string) (bool, error)
	InstallPackages(p []string, options ...string) (map[string]PkgInstalledResult, error)
}

func GetInstaller(name string) (PackageManager, error) {
	switch strings.ToLower(name) {
	case "brew":
		return Brew{}, nil
	case "apk":
		return nil, errors.New("Unsupported ATM :-(")
	case "apt":
		return nil, errors.New("Unsupported ATM :-(")
	case "chocolatey":
		return nil, errors.New("Unsupported ATM :-(")
	default:
		return nil, nil
	}
}
