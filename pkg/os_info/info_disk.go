package os_info

import (
	"code.cloudfoundry.org/bytefmt"
	"github.com/Benbentwo/SystemInfo/pkg/common/log"
	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/block"
	"github.com/jedib0t/go-pretty/v6/table"
	"strconv"
)

func getDisk() *ghw.BlockInfo {
	disk, err := ghw.Block()
	if err != nil {
		log.Logger().Errorf("Error getting block info: %v", err)
	}
	return disk
}

func (sysInfo *SystemInformation) outputDiskToTable() {
	if sysInfo.Disks != nil {

		diskCount := diskCountMap(sysInfo.Disks.Disks)

		SystemInfoWriter.AppendRow(table.Row{HEADER, "DISKs:", HEADER})

		for disk, count := range diskCount {
			SystemInfoWriter.AppendRow(table.Row{SPACE, disk + " x " + strconv.Itoa(count), SPACE, SPACE})
		}

		SystemInfoWriter.AppendRow(
			table.Row{"TOTAL: ", SPACE, SPACE, SPACE, SPACE, bytefmt.ByteSize(sysInfo.Disks.TotalPhysicalBytes)},
		)
		SystemInfoWriter.AppendSeparator()
	}
}

func diskCountMap(arr []*block.Disk) map[string]int {
	diskCountMap := make(map[string]int)
	for _, disk := range arr {
		diskCountMap[disk.DriveType.String()] = diskCountMap[disk.DriveType.String()] + 1
	}
	log.Logger().Debugf("Processor Map: %v", diskCountMap)
	return diskCountMap
}
