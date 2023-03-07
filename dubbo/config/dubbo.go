package config

import (
	"fmt"
)

type Config struct {
	//Application *Application `validate:"required" yaml:"application"`
	//Methods     map[string][]*Method        `validate:"required" yaml:"methods"`
	//Protocols   map[string]*Protocol        `validate:"required" yaml:"protocols"`
	Router []*Router `yaml:"router"`
	//Providers   []map[string]*Provider      `yaml:"providers"`
	//Names map[string]map[string]*Name `validate:"required" yaml:"names"`
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
