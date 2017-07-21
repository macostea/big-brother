package collector


import (
	"testing"
	"github.com/macostea/big-brother/model"
	"time"
	"encoding/json"
	"math"
)

func TestDataCollector_NewDataCollector(t *testing.T) {
	dataItems := []model.CollectedDataItem{&model.CPU{}}

	collector := NewDataCollector(dataItems)

	if len(collector.DataItems) != 1 {
		t.Fatal("Data collector does not have required number of data items")
	}

	if _, ok := collector.DataItems[0].(*model.CPU); ok == false {
		t.Fatal("Data collector data item is incorrect type")
	}
}

type MockDataType struct {
	Data map[string]string `json:"data"`
}

func (m *MockDataType) Type() string {
	return "mock"
}

func (m *MockDataType) GetInfo() {
	testInfo := make(map[string]string)
	testInfo["testKey"] = "testVal"

	m.Data = testInfo
}


func TestDataCollector_StartCollecting(t *testing.T) {
	dataItems := []model.CollectedDataItem{&MockDataType{}}

	collector := NewDataCollector(dataItems)

	startTime := time.Now()
	infoChannel := collector.StartCollecting(time.Second * 2)

	numberOfInfoMessages := 0
	for info := range infoChannel {
		if math.Abs(time.Since(startTime).Seconds()) - 2 > 0.01 {
			t.Fatal("Collector duration time not respected")
		}
		numberOfInfoMessages++

		var parsedInfo map[string]MockDataType

		err := json.Unmarshal(info, &parsedInfo)
		if err != nil {
			t.Fatal("Unable to parse collector info", err, info)
		}

		if parsedInfo["mock"].Data["testKey"] != "testVal" {
			t.Fatal("Received collector info is incorrect", parsedInfo)
		}

		if numberOfInfoMessages == 3 {
			break
		}

		startTime = time.Now()
	}
}
