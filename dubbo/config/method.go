package config

import "fmt"

type Method struct {
	Name    string `yaml:"name"`
	Retries string `yaml:"retries"`
}

func (m *Method) Prefix() string {
	return "dubbo.methods"
}

func (m *Method) Init(key string) {
	fmt.Println("dubbo method init", key)
}

func (m *Method) Order() int {
	return -1
}
