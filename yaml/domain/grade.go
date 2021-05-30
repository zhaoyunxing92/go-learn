package domain

type Grade struct {
	Id    string  `yaml:"id"`
	Name  string  `yaml:"name"`
	Users []*User `yaml:"users"`
}
