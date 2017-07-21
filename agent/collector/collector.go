package collector

import (
	"time"
	"github.com/mgutz/logxi/v1"
	"github.com/macostea/big-brother/model"
	"encoding/json"
)

type DataCollector struct {
	DataItems []model.CollectedDataItem
}

func NewDataCollector(dataItems []model.CollectedDataItem) *DataCollector {
	return &DataCollector{dataItems}
}

func (d *DataCollector) getAllInfo() map[string]model.CollectedDataItem {
	info := make(map[string]model.CollectedDataItem)
	for index := range d.DataItems {
		info[d.DataItems[index].Type()] = d.DataItems[index]
		d.DataItems[index].GetInfo()
	}

	return info
}

func (d *DataCollector) StartCollecting(duration time.Duration) <-chan []byte {
	infoChannel := make(chan []byte)
	ticker := time.NewTicker(duration)

	go func(channel chan<- []byte) {
		for range ticker.C {
			log.Debug("Time to get info")
			jsonString, err := json.Marshal(d.getAllInfo())
			if err != nil {
				log.Error("Failed to convert info to JSON", "err", err)
			}

			channel <- jsonString
		}
	}(infoChannel)

	return infoChannel
}
