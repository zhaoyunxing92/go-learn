package domain

type Application struct {
	Name        string `yaml:"name"`
	Module      string `yaml:"module"`
	Version     string `yaml:"version"`
	Owner       string `yaml:"owner"`
	Environment string `default:"dev" yaml:"environment"`
}
