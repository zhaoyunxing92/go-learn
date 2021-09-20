package service

//Config service
type Config struct {
	Id        string   `yaml:"id" json:"id"`
	Interface string   `yaml:"interface" json:"interface"`
	Registry  []string `yaml:"registry" json:"registry"`
}
