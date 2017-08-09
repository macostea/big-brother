package config

import "testing"

func TestServers_ReadConfig(t *testing.T) {
	masterConfig := MasterConfig{}

	masterConfig.ReadConfig("config_test.yaml")

	if len(masterConfig.Servers) != 1 {
		t.Fatal("Could not read servers array correctly")
	}

	server := masterConfig.Servers[0]

	if server.Port != "12345" {
		t.Fatal("Could not read server port correctly", "actual", server.Port, "expected", "12345")
	}

	if server.Name != "test" {
		t.Fatal("Could not read server name correctly", "actual", server.Name, "expected", "test")
	}

	if server.Addr != "test" {
		t.Fatal("Could not read server address correctly", "actual", server.Addr, "expected", "test")
	}
}