package config

import "fmt"

type Router struct {
	Tags []*Tag `yaml:"tags"`
}

func (r *Router) Init(key string) {
	fmt.Println("dubbo router init=", key)
}

func (r *Router) Order() int {
	return 4
}

type Tag struct {
	Name   string  `yaml:"name"`
	Remote *Remote `yaml:"remote"`
}

type Remote struct {
	Address []*Address `yaml:"address"`
}
type Address struct {
	Ip string
}

func (a Address) Init(key string) {

	fmt.Println("dubbo remote tag remote address init ", key)
}

func (r *Remote) Init(key string) {
	fmt.Println("dubbo remote init", key)
}

func (t *Tag) Init(key string) {
	fmt.Println("dubbo router tags init=", key)
}

func (t *Tag) Order() int {
	return 5
}
