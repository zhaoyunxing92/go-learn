package main

import "fmt"

/**
变量
* 标准声明
* 批量声明
* 变量类型推到和初始值
* 短变量声明（只能用在函数中）语法：:=
* 匿名变量（不占用命名空间，不会分配内存，可以重复使用，想忽略某个值的时候使用）语法：_
*/

func foo() (int, string) {
	return 10, "sunny"
}

func main() {
	//标准声明
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	// 批量声明
	var (
		a string
		b int
		c bool
		d int32
	)
	fmt.Println(a, b, c, d)
	// 变量类型推到和初始值
	var name1, age1 = "sunny", 28
	fmt.Println(name1, age1)

	// 短变量声明（只能用在函数中）
	e := 40
	fmt.Println(e)

	// 匿名变量
	x, _ := foo() //忽略第二个返回值
	_, y := foo() // 忽略第一个返回值

	fmt.Println(x, y)
}