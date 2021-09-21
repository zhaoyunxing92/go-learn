package domain

type Registry struct {
	Protocol string `yaml:"protocol"`
	Timeout  string `yaml:"timeout"`
	Group    string `yaml:"group"`
	Address  string `yaml:"address"`
}
