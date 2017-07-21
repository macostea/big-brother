package config

import (
	"testing"
	"time"
)

func TestAgentConfig_ReadConfig(t *testing.T) {
	ac := AgentConfig{}
	ac.ReadConfig("config_test.yaml")

	if ac.Server.Port != "1234" {
		t.Fatal("Config server port not read correctly")
	}

	if ac.Collector.Timeout != time.Second * 2 {
		t.Fatal("Config collector timeout not read correctly", ac.Collector.Timeout)
	}
}
