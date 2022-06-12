package main

import "fmt"

type Configuration struct {
	Broker struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"broker"`
	Stream struct {
		Deposits string `yaml:"deposits"`
	} `yaml:"stream"`
}

func (c Configuration) GetBroker() []string {
	return []string{fmt.Sprintf("%s:%s", c.Broker.Host, c.Broker.Port)}
}
