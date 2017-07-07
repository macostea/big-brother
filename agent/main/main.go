package main

import (
	"github.com/macostea/big-brother/agent/server"
	"time"
	"github.com/macostea/big-brother/agent/collector"
	"github.com/macostea/big-brother/agent/config"
)

func main() {
	c := config.AgentConfig{}
	c.ReadConfig("config.yaml") // TODO: pass this as a CLI parameter

	srv := server.AgentServer{}
	col := collector.DataCollector{}
	infoChannel := col.StartCollecting(time.Second * c.Collector.Timeout)

	go func(s *server.AgentServer) {
		for info := range infoChannel {
			s.SendInfoToClients(info)
		}
	}(&srv)

	srv.StartServer(c.Server.Port)
}
