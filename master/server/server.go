package server

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/mgutz/logxi/v1"
	"github.com/macostea/big-brother/model"
	"time"
	"container/list"
)

type S struct {
	upgrader websocket.Upgrader
	connections *list.List
}

func (s *S) statusHandler() func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		c, err := s.upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Error("Failed to create websocket", "err", err)
			return
		}

		s.connections.PushBack(c)
	}
}

func (s *S) SendInfoToClients(info map[string]map[string]model.CollectedDataItem) {
	toRemove := list.New()
	for e := s.connections.Front(); e != nil; e = e.Next() {
		if connection, ok := e.Value.(*websocket.Conn); ok {
			connection.SetReadDeadline(time.Now().Add(time.Millisecond * 10))

			if _, _, err := connection.NextReader(); err != nil {
				log.Debug("Websocket peer is not responding. Closing connection", "conn", connection, "err", err)
				connection.Close()
				toRemove.PushBack(e)
			}

			if err := connection.WriteJSON(info); err != nil {
				log.Error("Failed to write agent info", "err", err)
			}
		}
	}

	log.Debug("Cleaning up unused connections")
	for e := toRemove.Front(); e != nil; e = e.Next() {
		s.connections.Remove(e.Value.(*list.Element))
	}
}

func (s *S) StartServer() {
	s.upgrader = websocket.Upgrader{}

	s.connections = list.New()

	http.Handle("/", http.FileServer(http.Dir("master/client/build")))
	http.HandleFunc("/status", s.statusHandler())

	log.Fatal("Failed to start server", "err", http.ListenAndServe("localhost:8080", nil))
}
