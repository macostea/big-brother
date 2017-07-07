package config

import (
	"io/ioutil"
	"github.com/mgutz/logxi/v1"
	"gopkg.in/yaml.v2"
)

type Server struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}

type Servers struct {
	Servers []Server
}

func (s *Servers) ReadConfig(filename string) {
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Cannot read servers config file", "err", err)
	}

	yamlError := yaml.Unmarshal(configData, s)
	if yamlError != nil {
		log.Fatal("Cannot parse config file", "err", yamlError)
	}
}
