package server

import (
	"net"
	"github.com/mgutz/logxi/v1"
	"time"
	"container/list"
)

type AgentServer struct {
	connections *list.List
	isStarted bool
}

func (a *AgentServer) StartServer(port string) {
	ln, err := net.Listen("tcp", ":" + port)

	if err != nil {
		log.Fatal("Failed to start agent. Network error: ", "err", err)
		panic(err)
	}

	a.isStarted = true
	a.connections = list.New()

	for {
		if conn, err := ln.Accept(); err != nil {
			log.Error("Failed to accept connection: ", "err", err)
		} else {
			a.connections.PushBack(conn)
		}
	}
}

func (a *AgentServer) SendInfoToClients(info []byte) {
	buf := append(info, '\n')

	toRemove := list.New()

	for e := a.connections.Front(); e != nil; e = e.Next() {
		if connection, ok := e.Value.(net.Conn); ok {
			ack := [3]byte{}
			connection.SetReadDeadline(time.Now().Add(time.Millisecond * 10))
			if n, err := connection.Read(ack[0:3]); err != nil || n == 0 {
				log.Debug("Client disconnected", "conn", connection, "err", err)
				connection.Close()

				toRemove.PushBack(e)
				continue
			}

			if n, err := connection.Write(buf); err != nil || n == 0 {
				log.Error("Error writing to TCP connection", "err", err)
			}
		}
	}

	// Clean closed connections
	log.Debug("Cleaning closed connections")
	for e := toRemove.Front(); e != nil; e = e.Next() {
		a.connections.Remove(e.Value.(*list.Element))
	}

	log.Debug("Available connections", "connections", a.connections)
}
