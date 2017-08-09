package server

import (
	"github.com/macostea/big-brother/network"
	"github.com/mgutz/logxi/v1"
)

type AgentServer struct {
	N network.Net
	isStarted bool
}

func NewAgentServer(net network.Net) *AgentServer {
	return &AgentServer{net, false}
}

func (a *AgentServer) StartServer(port string) {
	go a.N.StartServer(port)
	a.isStarted = true
}

func (a *AgentServer) StopServer() {
	a.N.StopServer()

	a.isStarted = false
}

func (a *AgentServer) SendInfoToClients(info []byte) bool {
	if a.isStarted {
		a.N.SendToClients(info)
		return true
	} else {
		log.Error("Send info called when server is not running")
		return false
	}
}
