package placeholder

import (
	"fmt"

	"testing"
)

func TestResolver(t *testing.T) {
	str := New().
		Delims("#{", "}").
		Defaults(map[string]string{
			"username": "root",
			"password": "123",
		}).
		Resolver("java#{username}golang#{password:pwd}#{password:pwd}123#{password:pwd}")

	fmt.Println(str)
}
