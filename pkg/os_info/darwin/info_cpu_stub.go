// Package darwin
// +build !darwin

package darwin

import (
	"github.com/jaypipes/ghw"
)

func GetCpuDarwin() *ghw.CPUInfo {
	return nil
}
