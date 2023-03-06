package config

import "fmt"

type Application struct {
	Organization string    `default:"dubbo-go" yaml:"organization"`
	Name         string    `default:"dubbo.io" yaml:"name"`
	Module       string    `default:"sample" yaml:"module"`
	Group        string    `yaml:"group"`
	Version      string    `yaml:"version"`
	Owner        string    `default:"dubbo-go"`
	Environment  string    `yaml:"environment"`
	Metadata     *Metadata `default:"local" yaml:"metadata"`
}

func (a *Application) Prefix() string {
	return "dubbo.application"
}

func (a *Application) Init(key string) {
	fmt.Println("application init", key)
}

func (a *Application) Order() int {
	return 3
}
