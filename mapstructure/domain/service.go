package domain

type Service struct {
	Interface string `yaml:"interface"`
	Registry []string `yaml:"registry"`
}
