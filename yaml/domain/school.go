package domain

import "github.com/creasty/defaults"

// School 学校
type School struct {
	Name   string            `yaml:"name"`
	Grades map[string]*Grade `yaml:"grades"`
}

// Grade 班级
type Grade struct {
	Id    string  `yaml:"id"`
	Name  string  `yaml:"name"`
	Users []*User `yaml:"users"`
}

// User 用户
type User struct {
	Name   string `yaml:"name" default:"sunny"`
	Age    int    `yaml:"age" default:"20" json:"age,omitempty" property:"age"`
	Height int    `yaml:"height" default:"190" json:"height,omitempty" property:"height"`
}

// UnmarshalYAML 解析的时候给默认值
func (s *School) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(s); err != nil {
		return err
	}
	type plain School
	return unmarshal((*plain)(s))
}

func (u *User) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(u); err != nil {
		return err
	}
	type plain User
	return unmarshal((*plain)(u))
}

func (g *Grade) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(g); err != nil {
		return err
	}
	type plain Grade
	return unmarshal((*plain)(g))
}
