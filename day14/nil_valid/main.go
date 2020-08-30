package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a *int

	fmt.Println("var a *int is nil", reflect.ValueOf(a).IsNil()) //true
	fmt.Println("var a *int is valid", reflect.ValueOf(a).IsValid()) //true

    // 实例化一个匿名结构体
    b:= struct {}{}
    //反射获取其中字段
    fmt.Println("是否存在name字段",reflect.ValueOf(b).FieldByName("name").IsValid())
    fmt.Println("是否存在name方法",reflect.ValueOf(b).MethodByName("name").IsValid())
}
