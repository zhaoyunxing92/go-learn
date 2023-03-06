package config

import "fmt"

type Name struct {
	Address string `yaml:"address"`
}

func (n *Name) Prefix() string {
	return "dubbo.names"
}

func (n *Name) Init(key string) {

	fmt.Println("dubbo names init key=", key)
}
