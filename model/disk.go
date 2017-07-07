package model

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/mgutz/logxi/v1"
	"fmt"
	"time"
)

type Disk struct {
	Total uint64 `json:"total"`
	Free uint64 `json:"free"`
	Used uint64 `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	CollectionTime time.Time `json:"collection_time"`
}

func (d *Disk) GetInfo() {
	d.CollectionTime = time.Now()

	if stat, err := disk.Usage("/"); err != nil {
		log.Error("Failed to get disk info", err)
	} else {
		d.Total = stat.Total
		d.Free = stat.Free
		d.Used = stat.Used
		d.UsedPercent = stat.UsedPercent
	}
}

func (d *Disk) Type() string {
	return "disk"
}

func (d *Disk) String() string {
	return fmt.Sprintf("Total: %d, Free: %d, Used: %d, UsedPercent: %f", d.Total, d.Free, d.Used, d.UsedPercent)
}
