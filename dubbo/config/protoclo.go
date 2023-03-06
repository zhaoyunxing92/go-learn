package config

import "fmt"

type Protocol struct {
	Name   string      `default:"dubbo" yaml:"name"`
	Ip     string      `yaml:"ip"`
	Port   string      `default:"20000" yaml:"port"`
	Params interface{} `yaml:"params"`
}

func (p *Protocol) Prefix() string {
	return "dubbo.protocols"
}

func (p *Protocol) Init(key string) {
	if key == "dubbo.protocols.dubbo" {
		p.Name = "abc"
	}
	fmt.Println("protocol init", key)
}
