package model

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
	"github.com/mgutz/logxi/v1"
	"fmt"
)

type CPU struct {
	Count int `json:"count"`
	Load []float64 `json:"load"`
	CollectionTime time.Time `json:"collection_time"`
}

func (c *CPU) GetInfo() {
	c.CollectionTime = time.Now()

	if count, err := cpu.Counts(true); err != nil {
		log.Error("Failed to get cpu count", err)
	} else {
		c.Count = count
	}

	if load, err := cpu.Percent(time.Second * 2, true); err != nil {
		log.Error("Failed to get cpu load", err)
	} else {
		c.Load = load
	}
}

func (c *CPU) Type() string {
	return "cpu"
}

func (c *CPU) String() string {
	return fmt.Sprintf("Count: %d, Load: %+v", c.Count, c.Load)
}
