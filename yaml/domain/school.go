package domain

type School struct {
	Name   string            `yaml:"name"`
	Grades map[string]*Grade `yaml:"grades"`
}
