package main

import (
	"fmt"
	"reflect"
)

// 反射取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v)

	//获取值的类型
	k := v.Kind()
	switch k {

	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v,%T\n", ret, ret) //10,int32
	}
}

//set 值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v) //0xc00000a0c8,reflect.Value

	//Elem()根据指针取值
	k := v.Elem().Kind()
	switch k {

	case reflect.Int32:
		v.Elem().SetInt(50)
	}
}

func main() {

	var a int32 = 10

	reflectValue(a) //10,reflect.Value

	//使用指针
	reflectSetValue(&a)

	fmt.Println(a) //50
}
