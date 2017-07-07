package collector

import (
	"time"
	"github.com/mgutz/logxi/v1"
	"github.com/macostea/big-brother/model"
	"encoding/json"
)

var dataItems = []model.CollectedDataItem{&model.CPU{}, &model.Mem{}, &model.Disk{}}


type DataCollector struct {}

func (d *DataCollector) GetAllInfo() map[string]model.CollectedDataItem {
	info := make(map[string]model.CollectedDataItem)
	for index := range dataItems {
		info[dataItems[index].Type()] = dataItems[index]
		dataItems[index].GetInfo()
	}

	return info
}

func (d *DataCollector) StartCollecting(duration time.Duration) <-chan []byte {
	infoChannel := make(chan []byte)
	ticker := time.NewTicker(duration)

	go func(channel chan<- []byte) {
		for range ticker.C {
			log.Debug("Time to get info")
			jsonString, err := json.Marshal(d.GetAllInfo())
			if err != nil {
				log.Error("Failed to convert info to JSON", "err", err)
			}

			channel <- jsonString
		}
	}(infoChannel)

	return infoChannel
}
