package config

import "fmt"

type Router struct {
	Scope      string   `validate:"required" yaml:"scope"`
	Key        string   `validate:"required" yaml:"key"`
	Force      bool     `default:"false" yaml:"force"`
	Runtime    bool     `default:"false" yaml:"runtime"`
	Enabled    bool     `default:"true" yaml:"enabled"`
	Valid      bool     `default:"true" yaml:"valid"`
	Priority   int      `default:"0" yaml:"priority"`
	Conditions []string `yaml:"conditions"`
	Tags       []Tag    `yaml:"tags"`
}

func (r *Router) Prefix() string {
	return "dubbo.router"
}

func (r *Router) Init(key string) {
	fmt.Println("dubbo router init=", key)
}

func (r *Router) Order() int {
	return 4
}

type Tag struct {
	Name      string   `yaml:"name" json:"name,omitempty" property:"name"`
	Addresses []string `yaml:"addresses" json:"addresses,omitempty" property:"addresses"`
}

func (t *Tag) Prefix() string {
	return "dubbo.router.tags"
}

func (t *Tag) Init(key string) {
	fmt.Println("dubbo router tags init=", key)
}

func (t *Tag) Order() int {
	return 5
}
