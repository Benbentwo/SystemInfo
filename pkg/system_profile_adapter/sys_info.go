package system_profile_adapter

import (
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/utils"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/system_profiler"
	//"github.com/Benbentwo/system_profiler"
	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/pci"
	"github.com/pkg/errors"
)

func ExtractGraphicsInformation(sp *system_profiler.SystemProfiler) (*ghw.GPUInfo, error) {
	//log.Logger().Infoln("MEMORY: %s", sp.SPMemoryDataType)

	if sp == nil {
		return nil, errors.Errorf("System Profiler is %s, you should probably call .New()", utils.ColorBold("nil"))
	}
	// GRAPHICS
	//TODO right now i'm just doing a mapping of the field that I output.
	//  Really this should be an actual mapping, but to do this I need a
	//  Win Device so I can see how GHW maps out on windows and mirror
	//  on MacOSX
	graphics_cards := make([]*ghw.GraphicsCard, 0)
	for idx, display := range sp.SPDisplaysDataType {
		var vram = ""
		if display.SpdisplaysVram1 == "" {
			vram = display.SpdisplaysVram
		} else {
			vram = display.SpdisplaysVram1
		}
		graphicsCard := &ghw.GraphicsCard{
			Address: "",
			Index:   idx,
			DeviceInfo: &pci.Device{
				Revision: display.SppciModel,
				Address:  vram, // Yikes, really terrible mapping... again see above comment.
			},
			Node: nil,
		}
		graphics_cards = append(graphics_cards, graphicsCard)
	}
	return &ghw.GPUInfo{
		GraphicsCards: graphics_cards,
	}, nil
}
