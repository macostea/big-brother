package config

import (
	"io/ioutil"
	"github.com/mgutz/logxi/v1"
	"gopkg.in/yaml.v2"
	"time"
)

type CollectorConfig struct {
	Timeout time.Duration `yaml:"timeout"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type AgentConfig struct {
	Collector CollectorConfig `yaml:"collector"`
	Server ServerConfig `yaml:"server"`
}

func (ac *AgentConfig) ReadConfig(filename string) {
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Cannot read servers config file", "err", err)
	}

	yamlError := yaml.Unmarshal(configData, ac)
	if yamlError != nil {
		log.Fatal("Cannot parse config file", "err", yamlError)
	}
}
