package dubbo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhaoyunxing/dubbo/config"
	"github.com/zhaoyunxing/dubbo/extension"
)

func TestLoad(t *testing.T) {
	c := &config.Config{}

	if err := extension.Load("./dubbo/config", "dubbo", c); err != nil {
		t.Error(err)
	}
	definitions := extension.GetDefinitions()
	t.Log(len(definitions))
}

func TestMapToMap(t *testing.T) {
	names := make(map[string]map[string]*config.Name)
	names["tri"] = map[string]*config.Name{
		"json": {Address: "127.0.0.1"},
		"grpc": {Address: "127.0.0.1"},
	}
	names["dubbo"] = map[string]*config.Name{
		"json": {Address: "127.0.0.1"},
		"xml":  {Address: "127.0.0.1"},
	}
	conf := &config.Config{Names: names}
	extension.Register(conf)
	definitions := extension.GetDefinitions()
	assert.True(t, len(definitions) == 5)
}
