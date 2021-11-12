package os_info

import (
	"fmt"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/os_info/darwin"
	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/cpu"
	"github.com/jedib0t/go-pretty/v6/table"
	"runtime"
	"strconv"
)

func getCpu() *ghw.CPUInfo {
	if runtime.GOOS == "darwin" {
		return darwin.GetCpuDarwin()
	}
	cpuStats, err := ghw.CPU()
	if err != nil {
		log.Logger().Errorf("Error getting cpu info: %v", err)
	}
	return cpuStats
}

func (sysInfo *SystemInformation) outputCPUToTable() {
	if sysInfo.Cpu != nil {

		processorCount := processorCountMap(sysInfo.Cpu.Processors)
		SystemInfoWriter.AppendRow(table.Row{HEADER, "CPU:", HEADER})

		for processor, count := range processorCount {
			SystemInfoWriter.AppendRow(
				table.Row{SPACE, processor + " x " + strconv.Itoa(count), SPACE, SPACE},
			)
		}

		SystemInfoWriter.AppendRow(table.Row{"TOTAL", SPACE, SPACE, SPACE, fmt.Sprintf("Total Cores: %d", sysInfo.Cpu.TotalCores), fmt.Sprintf("Total Threads: %d", sysInfo.Cpu.TotalThreads)})
		SystemInfoWriter.AppendSeparator()
	}
}

func processorCountMap(arr []*cpu.Processor) map[string]int {
	processorCountMap := make(map[string]int)
	for _, processor := range arr {
		processorCountMap[processor.Model] = processorCountMap[processor.Model] + 1
	}
	log.Logger().Debugf("Processor Map: %s", processorCountMap)
	return processorCountMap
}
