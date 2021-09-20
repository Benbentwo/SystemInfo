// Package darwin
// +build !linux,!windows, darwin

package darwin

import (
	"github.com/jaypipes/ghw"
)

//import sysctl "github.com/lorenzosaino/go-sysctl"

func GetCpuDarwin() *ghw.CPUInfo {
	//_, err := cpu.New()
	//if err != nil {
	//	return nil
	//}
	//vals, err := sysctl.GetPattern("machdep.cpu.brand_string")
	//if err != nil {
	//	log.Logger().Errorln(err)
	//}
	return nil
}
