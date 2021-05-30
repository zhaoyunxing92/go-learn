package domain

type User struct {
	Name   string `yaml:"name" default:"sunny"`
	Age    int    `yaml:"age" default:"20" json:"age,omitempty" property:"age"`
	Height int    `yaml:"height" default:"190" json:"height,omitempty" property:"height"`
}
