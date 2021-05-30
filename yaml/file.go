package main

import (
	"fmt"
	"github.com/zhaoyunxing/yaml/domain"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	school := &domain.School{}
	bytes, err := ioutil.ReadFile(pwd + "\\yaml\\conf\\school.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.UnmarshalStrict(bytes, school)
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, grade := range school.Grades {
		for _, user := range grade.Users {
			fmt.Println("+++",key)
			fmt.Println("name:", user.Name)
			fmt.Println("age:", user.Age)
			fmt.Println("height:", user.Height)
		}
	}
}
