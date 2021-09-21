package domain

type DubboConfig struct {
	Application *Application         `yaml:"application"`
	Registries  map[string]*Registry `yaml:"registries"`
	Services    map[string]*Service  `yaml:"services"`
}
