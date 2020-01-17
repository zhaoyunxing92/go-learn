package main

import (
	"fmt"
)

// 匿名函数和闭包
func main() {
	// 赋值给一个变量
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()

	// 立即执行的匿名函数
	func() {
		fmt.Println("立即执行")
	}()
}