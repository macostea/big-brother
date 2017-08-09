package main

import (
	"github.com/macostea/big-brother/agent/server"
	"github.com/macostea/big-brother/agent/collector"
	"github.com/macostea/big-brother/agent/config"
	"flag"
	"github.com/macostea/big-brother/model"
	"github.com/macostea/big-brother/network"
)

func main() {
	var configFile = flag.String("config", "config.yaml", "Path to the config file")

	flag.Parse()

	c := config.AgentConfig{}
	c.ReadConfig(*configFile)

	networkController := network.NewController()
	srv := server.NewAgentServer(networkController)

	dataItems := []model.CollectedDataItem{&model.CPU{}, &model.Mem{}, &model.Disk{}}

	col := collector.NewDataCollector(dataItems)
	infoChannel := col.StartCollecting(c.Collector.Timeout)

	srv.StartServer(c.Server.Port)

	for info := range infoChannel {
		// Main app loop
		srv.SendInfoToClients(info)
	}
}
