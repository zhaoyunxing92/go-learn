package main

import "fmt"

var num = 10

/**
函数的变量访问规则
* 现在函数内部找，如果找到了则用，没有就找全局变量
*/
func testGlobal() {
	num = 20
	fmt.Println("全局变量：", num)
}

/**
求和
*/
func add(x, y int) int {
	return x + y
}

/**
差值
*/
func sub(x, y int) int {
	return x - y
}

/**
函数作为参数
*/
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	testGlobal()
	// 第三个参数是函数
	r1 := calc(100, 200, add)
	r2 := calc(100, 200, sub)
	fmt.Println(r1)
	fmt.Println(r2)
}
