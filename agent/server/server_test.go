package server

import (
	"testing"
	"time"
)

type MockNetwork struct {
	ServerStarted bool
	Done chan bool
	Port string
	InfoSent []byte
}

func (mn *MockNetwork) StartServer(port string) {
	mn.ServerStarted = true
	mn.Port = port
	mn.Done <- true
}

func (mn *MockNetwork) StopServer() {
	mn.ServerStarted = false
}

func (mn *MockNetwork) SendToClients(info []byte) {
	mn.InfoSent = info
}

func (mn *MockNetwork) ConnectToServer(address, port string) {
	// Not used here
}

func (mn *MockNetwork) StartReadingFromServer(infoChanel chan <-[]byte) {
	// Not used here
}

func TestAgentServer_StartServer(t *testing.T) {
	mn := &MockNetwork{}
	mn.Done = make(chan bool)

	s := NewAgentServer(mn)

	s.StartServer("1234")

	select {
	case <-mn.Done:
	case <-time.After(2 * time.Second):
		t.Fatal("Timed out waiting for server to start")
	}

	if s.isStarted != true {
		t.Fatal("Server did not start correctly")
	}

	if mn.ServerStarted != true {
		t.Fatal("Server did not call the network controller correctly on start")
	}

	if mn.Port != "1234" {
		t.Fatal("Server did not pass the correct port to the network controller", "actual", mn.Port, "expected", "1234")
	}
}

func TestAgentServer_StopServer(t *testing.T) {
	mn := &MockNetwork{}
	mn.Done = make(chan bool)
	s := NewAgentServer(mn)

	s.StartServer("1234")

	select {
	case <-mn.Done:
	case <-time.After(2 * time.Second):
		t.Fatal("Timed out waiting for server to start")
	}

	s.StopServer()

	if s.isStarted {
		t.Fatal("Server did not stop correctly")
	}

	if mn.ServerStarted {
		t.Fatal("Server did not call the network controller correctly on stop")
	}
}

func TestAgentServer_SendInfoToClients(t *testing.T) {
	mn := &MockNetwork{}
	mn.Done = make(chan bool)

	s := NewAgentServer(mn)

	s.StartServer("1234")

	select {
	case <-mn.Done:
	case <-time.After(2 * time.Second):
		t.Fatal("Timed out waiting for server to start")
	}

	message := []byte("test")

	success := s.SendInfoToClients(message)
	if !success || string(mn.InfoSent) != string(message) {
		t.Fatal("Server did not send message correctly", "actual", mn.InfoSent, "expected", message)
	}

	s.StopServer()
	success = s.SendInfoToClients(message)
	if success {
		t.Fatal("Server tried to send message to clients even if it is not running")
	}
}
