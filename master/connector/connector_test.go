package connector

import (
	"testing"
	"github.com/macostea/big-brother/master/config"
	"time"
	"github.com/macostea/big-brother/model"
	"encoding/json"
	"github.com/mgutz/logxi/v1"
)

type MockNetwork struct {
	connected bool
	connectionAddress string
	connectionPort string
}

func (mn *MockNetwork) StartServer(port string) {
	// Not used here
}

func (mn *MockNetwork) StopServer() {
	// Not used here
}

func (mn *MockNetwork) SendToClients(info []byte) {
	// Not used here
}

func (mn *MockNetwork) ConnectToServer(address, port string) {
	mn.connected = true
	mn.connectionAddress = address
	mn.connectionPort = port
}

func (mn *MockNetwork) StartReadingFromServer(infoChanel chan <-[]byte) {
	ticker := time.NewTicker(time.Second * 1)

	n := 0

	go func() {
		for range ticker.C {
			if n == 2 {
				close(infoChanel)
				ticker.Stop()
				break
			}

			info := map[string]model.CollectedDataItem{}
			info["cpu"] = &model.CPU{Count: n, Load: nil, CollectionTime: time.Now()}

			data, err := json.Marshal(info)
			if err != nil {
				log.Fatal("Failed to convert test data to json")
			}

			infoChanel <- data
			n++
		}
	}()
}

func TestC_ConnectToServer(t *testing.T) {
	mn := &MockNetwork{}

	serverConnector := NewServerConnector(mn)

	server := config.Server{"test", "test", "12345"}
	serverConnector.ConnectToServer(server)

	if mn.connectionPort != "12345" {
		t.Fatal("Connector did not set up the connection port correctly", "actual", mn.connectionPort, "expected", "12345")
	}

	if mn.connectionAddress != "test" {
		t.Fatal("Connector did not set up the connection address correctly", "actual", mn.connectionAddress, "expected", "test")
	}

	if !mn.connected {
		t.Fatal("Connector did not call the network controller connection")
	}

	n := 0
	for info := range serverConnector.InfoChannel {
		cpu := info["test"]["cpu"].(*model.CPU)
		if cpu.Count != n {
			t.Fatal("Connector did not send correct data")
		}

		n++
	}
}