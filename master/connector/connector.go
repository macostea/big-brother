package connector

import (
	"github.com/macostea/big-brother/master/config"
	"net"
	"github.com/mgutz/logxi/v1"
	"bufio"
	"fmt"
	"github.com/macostea/big-brother/model"
	"encoding/json"
)

type C struct {
	conn net.Conn
	InfoChannel <-chan map[string]map[string]model.CollectedDataItem
}

func (mc *C) ConnectToServer(server config.Server) {
	conn, err := net.Dial("tcp", server.Addr + ":" + server.Port)

	if err != nil {
		log.Error("Failed to connect to server", "srvname", server.Name, "err", err)
		return
	}

	mc.conn = conn
	infoChannel := make(chan map[string]map[string]model.CollectedDataItem)
	mc.InfoChannel = infoChannel

	go handleConnection(conn, server, infoChannel)
}

func handleConnection(conn net.Conn, server config.Server, infoChannel chan<- map[string]map[string]model.CollectedDataItem) {
	fmt.Fprintf(conn, "%s", "ack")

	defer conn.Close()
	for {
		status, err := bufio.NewReader(conn).ReadBytes('\n')

		if err != nil {
			log.Error("Failed to read from server", "err", err)
			break
		}

		var info map[string]*json.RawMessage

		err = json.Unmarshal(status, &info)
		if err != nil {
			log.Error("Failed to parse JSON from server", "err", err)
			fmt.Printf("%v", string(status))
			fmt.Printf("%v", err)
			break
		}

		finalInfo := map[string]map[string]model.CollectedDataItem{}

		finalInfo[server.Name] = map[string]model.CollectedDataItem{}

		for key, val := range info {
			switch key {
			case "cpu":
				obj := &model.CPU{}
				json.Unmarshal(*val, obj)
				finalInfo[server.Name][key] = obj
			case "disk":
				obj := &model.Disk{}
				json.Unmarshal(*val, obj)
				finalInfo[server.Name][key] = obj
			case "mem":
				obj := &model.Mem{}
				json.Unmarshal(*val, obj)
				finalInfo[server.Name][key] = obj
			}
		}

		infoChannel <- finalInfo

		fmt.Fprintf(conn, "%s", "ack")
	}
}


