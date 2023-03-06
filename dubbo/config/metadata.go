package config

import "fmt"

type Metadata struct {
	Type        string         `json:"type"`
	MetadataTag []*MetadataTag `json:"tags"`
}
type MetadataTag struct {
	Name string `json:"name"`
}

func (m *Metadata) Prefix() string {
	return "dubbo.application.metadata"
}

func (m *Metadata) Init(key string) {
	fmt.Println("dubbo metadata init", key)
}
