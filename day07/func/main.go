package main

import (
	"fmt"
	"strings"
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

	// 闭包
	r := makeSuffixFunc(".txt")
	ret := r("zhaoyunxing")
	fmt.Println(ret)
	//
	x, y := calc(100)
	ret1 := x(200) //base= 100+200 300
	ret2 := y(100) //base= 300-100 200
	fmt.Println(ret1, ret2)
}

/**
判断字符串结尾,如果是则返回，如果不是则拼接结尾后返回
*/
func makeSuffixFunc(suffix string) func(string) string {
	return func(str string) string {
		//判断字符串的结尾是否符合要求
		if !strings.HasSuffix(str, suffix) {
			return str + suffix
		}
		return str
	}
}

/**
变量复用
*/
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
