package main

import "fmt"

//1.自定义类型
type MyInt int

// 定义别名
type NewInt = int

//类型别名和自定义类型
func main() {

	var a MyInt
	fmt.Printf("type:%T value:%v\n", a, a) //type:main.MyInt value:0

	var b NewInt
	fmt.Printf("type:%T value:%v\n", b, b) //type:int value:0
}
