package registry

type Config struct {
	Protocol string `yaml:"protocol" json:"protocol"`
	Timeout  string `yaml:"timeout" json:"timeout"`
	Group    string `yaml:"group" json:"group"`
	Address  string `yaml:"address" json:"address"`
}