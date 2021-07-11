package os_info

import (
	"code.cloudfoundry.org/bytefmt"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/jaypipes/ghw"
	"github.com/jedib0t/go-pretty/v6/table"
)

func getDisk() *ghw.BlockInfo {
	disk, err := ghw.Block()
	if err != nil {
		log.Logger().Errorf("Error getting block info: %v", err)
	}
	return disk
}

func (sysInfo *SystemInformation) outputDiskToTable() {
	sysInfo.Table.AppendHeader(table.Row{HEADER, "DISKs:", HEADER})

	sysInfo.Table.AppendRow(
		table.Row{"TOTAL: ", SPACE, SPACE, SPACE, SPACE, bytefmt.ByteSize(sysInfo.Disks.TotalPhysicalBytes)},
	)
	for _, disk := range sysInfo.Disks.Disks {
		sysInfo.Table.AppendRow(
			table.Row{SPACE, disk.Name, disk.Vendor, bytefmt.ByteSize(disk.SizeBytes)},
		)
		for _, partition := range disk.Partitions {
			sysInfo.Table.AppendRow(
				table.Row{SPACE, SPACE, partition.Name, partition.Label, partition.Type, bytefmt.ByteSize(partition.SizeBytes)},
			)
		}
	}
	sysInfo.Table.AppendSeparator()
}
