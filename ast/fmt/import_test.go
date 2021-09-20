package fmt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess(t *testing.T) {

	str := `package app
import "time"

// import test
import (
  // abc
  "fmt"
)

import "time"

// abc
var (
  a int
  b int
)

const (
	z = 10
	q = "dubbo"
)
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
`
	out, err := Process("", []byte(str))

	assert.Nil(t, err)
	fmt.Println(string(out))
}
