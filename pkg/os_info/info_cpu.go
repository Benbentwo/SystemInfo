package os_info

import (
	"fmt"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/cpu"
	"github.com/jedib0t/go-pretty/v6/table"
	"strconv"
)

func getCpu() *ghw.CPUInfo {
	cpu, err := ghw.CPU()
	if err != nil {
		log.Logger().Errorf("Error getting cpu info: %v", err)
	}
	return cpu
}

func (sysInfo *SystemInformation) outputCPUToTable() {
	processorCount := processorCountMap(sysInfo.Cpu.Processors)
	for processor, count := range processorCount {
		sysInfo.Table.AppendHeader(table.Row{HEADER, "CPU:", HEADER})
		sysInfo.Table.AppendRow(
			table.Row{SPACE, processor + " x " + strconv.Itoa(count), fmt.Sprintf("Total Cores: %d", sysInfo.Cpu.TotalCores), fmt.Sprintf("Total Threads: %d", sysInfo.Cpu.TotalThreads)},
		)
	}
	sysInfo.Table.AppendSeparator()
}

// I don't think a single computer should have different CPUs and different processors, but just incase.
func processorCountMap(arr []*cpu.Processor) map[string]int {
	processorCountMap := make(map[string]int)
	for _, processor := range arr {
		processorCountMap[processor.Model] = processorCountMap[processor.Model] + 1
	}
	log.Logger().Debugf("Processor Map: %s", processorCountMap)
	return processorCountMap
}
