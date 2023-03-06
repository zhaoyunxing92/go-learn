package config

import (
	"fmt"
)

type Config struct {
	//Methods     map[string][]*Method `validate:"required" yaml:"methods"`
	//Application *Application `validate:"required" yaml:"application"`
	//Protocols   map[string]*Protocol `validate:"required" yaml:"protocols"`
	//Router      []*Router            `yaml:"router"`
	//Name        string               `validate:"required" yaml:"string"`
	Provider []map[string]*Provider `yaml:"providers"`
}

func (c *Config) Prefix() string {
	return "dubbo"
}

func (c *Config) Init(key string) {
	fmt.Println("config dubbo init", key)
}

func (c *Config) Order() int {
	return 0
}
