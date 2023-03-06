package config

import "fmt"

type Provider struct {
	Register bool   `yaml:"register"`
	Name     string `yaml:"name"`
}

func (p *Provider) Prefix() string {
	return "dubbo.providers"
}

func (p *Provider) Init(key string) {
	fmt.Println("provider init", key)
}
