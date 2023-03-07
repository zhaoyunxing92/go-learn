package dubbo

import (
	"github.com/stretchr/testify/assert"
	"testing"

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

func TestSimpleStruct(t *testing.T) {
	c := &config.Config{Application: &config.Application{Organization: "dubbo"}}
	extension.Register(c)
	definitions := extension.GetKeys()

	keys := []string{"dubbo", "dubbo.application"}
	assert.True(t, len(definitions) == len(keys))
	for _, key := range keys {
		if _, ok := definitions[key]; !ok {
			t.Error("key def not find", key)
		}
	}
}

func TestMapAndStruct(t *testing.T) {
	c := &config.Config{Protocols: map[string]*config.Protocol{
		"dubbo": {
			Name: "dubbo",
			Ip:   "127.0.0.1",
			Params: []*config.Param{
				{Key: "username", Val: "nacos"},
				{Key: "pwd", Val: "123456"}},

			Tags: map[string]*config.Tag{
				"user": {Name: "host", Remote: &config.Remote{Address: []*config.Address{
					{Ip: "127.0.0.1"},
				}}},
				"order": {Name: "host", Remote: &config.Remote{}},
			},
		},
		"tri": {Name: "tri", Ip: "127.0.0.1"},
	}}

	extension.Register(c)
	definitions := extension.GetKeys()

	keys := []string{
		"dubbo",
		"dubbo.dubbo",
		"dubbo.dubbo.user",
		"dubbo.dubbo.user.remote",
		"dubbo.dubbo.user.remote.address.0",
		"dubbo.dubbo.params.1",
		"dubbo.dubbo.params.0",
		"dubbo.tri",
		"dubbo.dubbo.order",
		"dubbo.dubbo.order.remote",
	}
	assert.True(t, len(definitions) == len(keys))
	for _, key := range keys {
		if _, ok := definitions[key]; !ok {
			t.Error("key def not find", key)
		}
	}
}

// map[string]map[string]*Name
func TestMapAndMap(t *testing.T) {
	names := make(map[string]map[string]*config.Name)
	names["tri"] = map[string]*config.Name{
		"json": {Address: "127.0.0.1"},
		"grpc": {Address: "127.0.0.1"},
	}
	names["dubbo"] = map[string]*config.Name{
		"json": {Address: "127.0.0.1"},
		"xml":  {Address: "127.0.0.1"},
	}

	keys := []string{"dubbo", "dubbo.names.tri.grpc", "dubbo.names.tri.json",
		"dubbo.names.dubbo.json", "dubbo.names.dubbo.xml"}

	conf := &config.Config{Names: names}
	extension.Register(conf)

	definitions := extension.GetKeys()
	assert.True(t, len(definitions) == len(keys))
	for _, key := range keys {
		if _, ok := definitions[key]; !ok {
			t.Error("key def not find", key)
		}
	}
}

// [] -> [] -> []
func TestSliceAndSliceAndStruct(t *testing.T) {
	router := make([]*config.Router, 0, 3)
	address := make([]*config.Address, 0, 3)
	address = append(address, &config.Address{Ip: "127.0.0.1"})

	router = append(router, &config.Router{Tags: []*config.Tag{
		{Name: "dubbo", Remote: &config.Remote{Address: address}},
		{Name: "tri"},
	}})

	router = append(router, &config.Router{Tags: []*config.Tag{
		{Name: "tri"},
	}})

	c := &config.Config{Router: router}

	extension.Register(c)
	keys := []string{
		"dubbo",
		"dubbo.router.0",
		"dubbo.router.0.tags.0",
		"dubbo.router.0.tags.0.remote",
		"dubbo.router.0.tags.0.remote.address.0",
		"dubbo.router.0.tags.1",
		"dubbo.router.1",
		"dubbo.router.1.tags.0",
	}
	definitions := extension.GetKeys()
	assert.True(t, len(definitions) == len(keys))
	for _, key := range keys {
		if _, ok := definitions[key]; !ok {
			t.Error("key def not find", key)
		}
	}
}
