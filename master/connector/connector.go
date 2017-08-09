package connector

import (
	"github.com/macostea/big-brother/master/config"
	"net"
	"github.com/mgutz/logxi/v1"
	"fmt"
	"github.com/macostea/big-brother/model"
	"encoding/json"
	"github.com/macostea/big-brother/network"
)

type C struct {
	conn net.Conn
	InfoChannel <-chan map[string]map[string]model.CollectedDataItem
	privateInfoChannel chan map[string]map[string]model.CollectedDataItem
	N network.Net
}

func NewServerConnector(net network.Net) *C {
	infoChannel := make(chan map[string]map[string]model.CollectedDataItem)
	return &C{nil, infoChannel, infoChannel, net}
}

func (mc *C) ConnectToServer(server config.Server) {
	mc.N.ConnectToServer(server.Addr, server.Port)

	infoChannel := make(chan []byte)

	go mc.N.StartReadingFromServer(infoChannel)

	go func() {
		for networkInfo := range infoChannel {
			finalInfo := map[string]map[string]model.CollectedDataItem{}
			convertedInfo := convertNetworkInfo(networkInfo)

			if convertedInfo == nil {
				break // Error should already be logged
			}

			finalInfo[server.Name] = convertedInfo
			mc.privateInfoChannel <- finalInfo
		}

		close(mc.privateInfoChannel)
	}()
}

func convertNetworkInfo(networkInfo []byte) map[string]model.CollectedDataItem {
	var info map[string]*json.RawMessage

	err := json.Unmarshal(networkInfo, &info)
	if err != nil {
		log.Error("Failed to parse JSON from server", "err", err)
		fmt.Printf("%v", string(networkInfo))
		fmt.Printf("%v", err)

		return nil
	}

	serverInfo := map[string]model.CollectedDataItem{}

	// TODO: Maybe remove this code. It is redundant now that we have the type
	for key, val := range info {
		var obj model.CollectedDataItem
		switch key {
		case "cpu":
			obj = &model.CPU{}
		case "disk":
			obj = &model.Disk{}
		case "mem":
			obj = &model.Mem{}
		}

		json.Unmarshal(*val, obj)
		serverInfo[key] = obj
	}

	return serverInfo
}
