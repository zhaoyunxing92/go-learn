package main

import (
	"fmt"
	"github.com/zhaoyunxing/dubbo/config"
	"github.com/zhaoyunxing/dubbo/extension"
)

func main() {

	c := &config.Config{}

	if err := extension.Load("./dubbo/config", "dubbo", c); err != nil {
		panic(err)
	}

	definitions := extension.GetDefinitions()
	fmt.Println(len(definitions))
}
