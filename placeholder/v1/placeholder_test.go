package v1

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	str := New().
		Delims("#{", "}").
		Resolver("keep#{name:zyx}#{age:30}")

	fmt.Println(str)
}
