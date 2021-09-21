package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYamlLoad(t *testing.T) {
	config := Load(WithPath("../conf/yaml"))

	application, _ := config.GetApplicationConfig()

	assert.Equal(t, application.Name, "dubbo-go")
}

func TestPropertiesLoad(t *testing.T) {
	conf := Load(WithPath("../conf/properties"),
		WithName("application.properties"),
		WithGenre("properties"))

	application, _ := conf.GetApplicationConfig()

	assert.Equal(t, application.Name, "dubbo-go")
}
