package fmt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess(t *testing.T) {

	str := `package app
import "time"

import(
    "dubbo.apache.org/dubbo-go/v3/protocol"
    "github.com/creasty/defaults"
)

import "time"

// Config is a configuration for current application, whether the application is a provider or a consumer
type Config struct {
	Organization string 
	Name         string 
	Module       string 
	Version      string 
	Owner        string 
	Environment  string 
	// the metadata type. remote or local
	MetadataType string 
}

// Prefix dubbo.application
func (Config) Prefix() string {
	return "dubbo" + ".application"
}

func (c *Config) SetDefault() error {
	return defaults.Set(c)
}

func (c *Config) Validate(valid *validator.Validate) error {
	if err := valid.Struct(c); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Error())
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}

// UnmarshalYAML unmarshal the Config by @unmarshal function
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(c); err != nil {
		return err
	}
	type plain Config
	return unmarshal((*plain)(c))
}
`
	_, err := Process("", []byte(str))
	assert.Nil(t, err)
	//fmt.Println(string(out))
}
