package server

import (
	"github.com/macostea/big-brother/network"
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

func (a *AgentServer) SendInfoToClients(info []byte) {
	a.N.SendToClients(info)
}
