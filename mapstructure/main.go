package main

import (
	"github.com/mitchellh/mapstructure"
	"github.com/zhaoyunxing/mapstructure/domain"
	"log"
)

func main() {

	m := map[string]interface{}{
		"dubbo.application.name":    "dubbo-go",
		"dubbo.application.module":  "local",
		"dubbo.application.version": "1.0.0",
		"dubbo.application.owner":   "zhaoyunxing",
	}

	var conf domain.DubboConfig
	var metadata mapstructure.Metadata

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &conf,
		Metadata:         &metadata,
		TagName:          "yaml",
	})

	if err != nil {
		log.Fatal(err)
	}

	err = decoder.Decode(m)
	if err == nil {

	} else {
		log.Println(err.Error())
	}
}
