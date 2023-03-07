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

// map[string]map[string]*Name
func TestMapAndMap(t *testing.T) {
	//names := make(map[string]map[string]*config.Name)
	//names["tri"] = map[string]*config.Name{
	//	"json": {Address: "127.0.0.1"},
	//	"grpc": {Address: "127.0.0.1"},
	//}
	//names["dubbo"] = map[string]*config.Name{
	//	"json": {Address: "127.0.0.1"},
	//	"xml":  {Address: "127.0.0.1"},
	//}
	//
	//keys := []string{"dubbo", "dubbo.names.tri.grpc", "dubbo.names.tri.json",
	//	"dubbo.names.dubbo.json", "dubbo.names.dubbo.xml"}
	//
	//conf := &config.Config{Names: names}
	//extension.Register(conf)
	//
	//definitions := extension.GetKeys()
	//assert.True(t, len(definitions) == len(keys))
	//for _, key := range keys {
	//	if _, ok := definitions[key]; !ok {
	//		t.Error("key def not find", key)
	//	}
	//}
}

func TestSlice(t *testing.T) {
	router := make([]*config.Router, 0, 3)
	router = append(router, &config.Router{Scope: "local", Tags: []config.Tag{
		{Name: "dubbo", Addresses: "127.0.0.1"},
	}})

	router = append(router, &config.Router{Scope: "remote", Tags: []config.Tag{
		{Name: "tri", Addresses: "127.0.0.1"},
	}})

	c := &config.Config{Router: router}

	extension.Register(c)

	definitions := extension.GetKeys()
	assert.True(t, len(definitions) > 0)
}
