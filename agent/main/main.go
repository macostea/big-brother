package main

import (
	"github.com/macostea/big-brother/agent/server"
	"time"
	"github.com/macostea/big-brother/agent/collector"
	"github.com/macostea/big-brother/agent/config"
	"flag"
	"github.com/macostea/big-brother/model"
)

func main() {
	var configFile = flag.String("config", "config.yaml", "Path to the config file")

	flag.Parse()

	c := config.AgentConfig{}
	c.ReadConfig(*configFile)

	srv := server.AgentServer{}

	dataItems := []model.CollectedDataItem{&model.CPU{}, &model.Mem{}, &model.Disk{}}

	col := collector.NewDataCollector(dataItems)
	infoChannel := col.StartCollecting(time.Second * c.Collector.Timeout)

	go func(s *server.AgentServer) {
		for info := range infoChannel {
			s.SendInfoToClients(info)
		}
	}(&srv)

	srv.StartServer(c.Server.Port)
}
