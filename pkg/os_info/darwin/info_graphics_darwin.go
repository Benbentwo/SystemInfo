package darwin

import (
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/system_profile_adapter"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/system_profiler"
	//"github.com/Benbentwo/system_profiler"
	"github.com/jaypipes/ghw"
)

func GetGpuDarwin() (*ghw.GPUInfo, error) {
	system_profile, err := system_profiler.New()
	if err != nil {
		log.Logger().Errorf("fetching system_profile failed: %s", err)
		return nil, err
	}
	return system_profile_adapter.ExtractGraphicsInformation(system_profile)
}
