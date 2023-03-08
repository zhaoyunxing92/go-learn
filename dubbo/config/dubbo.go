package config

import (
	"fmt"
)

type Config struct {
	Protocols   map[string]*Protocol        `validate:"required" yaml:"protocols"`
	Names       map[string]map[string]*Name `validate:"required" yaml:"names"`
	Application *Application                `validate:"required" yaml:"application"`
	Methods     map[string][]*Method        `validate:"required" yaml:"methods"`
	Router      []*Router                   `yaml:"router"`
	Providers   []map[string]*Provider      `yaml:"providers"`
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
