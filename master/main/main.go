package main

import (
	"github.com/macostea/big-brother/master/config"
	"github.com/macostea/big-brother/master/connector"
	"github.com/macostea/big-brother/master/server"
)

func main() {
	servers := config.Servers{}
	servers.ReadConfig("master/config/servers.yaml")

	c := connector.C{}
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
