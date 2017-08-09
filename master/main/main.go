package main

import (
	"github.com/macostea/big-brother/master/config"
	"github.com/macostea/big-brother/master/connector"
	"github.com/macostea/big-brother/master/server"
	"flag"
	"github.com/macostea/big-brother/network"
)

func main() {
	var configFile = flag.String("config", "config.yaml", "Path to the config file")

	flag.Parse()

	servers := config.MasterConfig{}
	servers.ReadConfig(*configFile)

	networkController := network.NewController()

	c := connector.NewServerConnector(networkController)
	for _, s := range servers.Servers {
		c.ConnectToServer(s)
	}

	masterServer := server.S{}

	go func(s *server.S) {
		for info := range c.InfoChannel {
			s.SendInfoToClients(info)
		}
	}(&masterServer)

	masterServer.StartServer()

}
