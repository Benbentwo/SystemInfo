package package_managers

import (
	"github.com/Benbentwo/SystemInfo/pkg/common/log"
	"os/exec"
)

func Execute(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.Output()

	if err != nil {
		log.Logger().Errorf(err.Error())
		return string(stdout), err
	}

	// Print the output
	return string(stdout), nil
}
