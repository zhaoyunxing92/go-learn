package config

import "fmt"

type Protocol struct {
	Name   string          `default:"dubbo" yaml:"name"`
	Ip     string          `yaml:"ip"`
	Port   string          `default:"20000" yaml:"port"`
	Params []*Param        `yaml:"params"`
	Tags   map[string]*Tag `validate:"required" yaml:"tags"`
}

type Param struct {
	Key string
	Val string
}

func (p *Param) Init(key string) {
	fmt.Println("protocol param init", key)
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
