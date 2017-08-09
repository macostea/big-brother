package network

import (
	"net"
	"github.com/mgutz/logxi/v1"
	"container/list"
	"time"
)

type Net interface {
	StartServer(port string)
	StopServer()
	SendToClients(info []byte)
}

type Controller struct {
	Connections *list.List
	listener net.Listener
}

func NewController() *Controller {
	return &Controller{list.New(), nil}
}

func (nc *Controller) StartServer(port string) {
	ln, err := net.Listen("tcp", ":" + port)

	if err != nil {
		log.Fatal("Failed to start agent. Network error: ", "err", err)
	}

	for {
		if conn, err := ln.Accept(); err != nil {
			log.Error("Failed to accept connection: ", "err", err)
		} else {
			nc.Connections.PushBack(conn)
		}
	}
}

func (nc *Controller) StopServer() {
	err := nc.listener.Close()

	if err != nil {
		log.Error("Failed to stop server", "err", err)
	}
}

func (nc *Controller) SendToClients(info []byte) {
	connectionsToRemove := list.New()

	for e := nc.Connections.Front(); e != nil; e = e.Next() {
		if connection, ok := e.Value.(net.Conn); ok {
			if !isClientAlive(connection) {
				connectionsToRemove.PushBack(e)
				continue
			}

			if n, err := connection.Write(info); err != nil || n == 0 {
				log.Error("Error writing to TCP connection", "err", err)
				connectionsToRemove.PushBack(e)
				connection.Close()
			}
		}
	}

	nc.cleanConnections(connectionsToRemove)
}

func isClientAlive(client net.Conn) bool {
	ack := [3]byte{}
	client.SetReadDeadline(time.Now().Add(time.Millisecond * 30))

	if n, err := client.Read(ack[0:3]); err != nil || n == 0 {
		log.Debug("Client disconnected", "conn", client, "err", err)
		client.Close()

		return false
	}

	return true
}

func (nc *Controller) cleanConnections(connectionsToRemove *list.List) {
	log.Debug("Cleaning closed connections")
	for e := connectionsToRemove.Front(); e != nil; e = e.Next() {
		nc.Connections.Remove(e.Value.(*list.Element))
	}
}