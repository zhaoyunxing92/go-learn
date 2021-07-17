package annotation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAnnotation(t *testing.T) {
	str := `@DubboService(interfaceName = "helloService",interface="com.cloud.dubbo.service.HelloService", version = "1.0.0", group = "dubbo")`

	var (
		ant Annotation
		err error
	)

	ant, err = ParseAnnotation(str)

	assert.Nil(t, err)
	assert.Equal(t, "DubboService", ant.Name)
	assert.Equal(t, "helloService", ant.Attributes["interfaceName"])
	assert.Equal(t, "com.cloud.dubbo.service.HelloService", ant.Attributes["interface"])
	assert.Equal(t, "1.0.0", ant.Attributes["version"])
	assert.Equal(t, "dubbo", ant.Attributes["group"])
}
