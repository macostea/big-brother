package model

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/mgutz/logxi/v1"
	"fmt"
	"time"
)

type Mem struct {
	Total uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used uint64 `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Free uint64 `json:"free"`
	CollectionTime time.Time `json:"collection_time"`
}

func (m *Mem) GetInfo() {
	m.CollectionTime = time.Now()

	if stat, err := mem.VirtualMemory(); err != nil {
		log.Error("Failed to get memory info", err)
	} else {
		m.Total = stat.Total
		m.Available = stat.Available
		m.Used = stat.Used
		m.UsedPercent = stat.UsedPercent
		m.Free = stat.Free
	}
}

func (m *Mem) Type() string {
	return "mem"
}

func (m *Mem) String() string {
	return fmt.Sprintf("Total: %d, Available: %d, Used: %d, UsedPercent: %f, Free: %d", m.Total, m.Available, m.Used, m.UsedPercent, m.Free)
}
