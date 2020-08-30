package main

import (
	"fmt"
	"reflect"
)

// 反射
func reflectType(x interface{}) {

	obj := reflect.TypeOf(x)
	fmt.Printf("type:%v,kind:%v\n", obj.Name(), obj.Kind())
}



type Cat struct{}

type Dog struct{}

func main() {

	var a float32 = 3.1415926

	reflectType(a) // type:float32,kind:float32

	var b int8 = 10

	reflectType(b) // type:int8,kind:int8

	var cat Cat

	reflectType(cat) //type:Cat,kind:struct

	var dog Dog

	reflectType(dog) //type:Dog,kind:struct

	// 切片、Map、指针、数组他们的.Name()都是空
	var slice []int

	reflectType(slice) //type:,kind:slice

	var str []string

	reflectType(str) //type:,kind:slice

}
