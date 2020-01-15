package main

import "fmt"

// for 循环结构
func main() {
	//1. 基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2.省略初始语句
	fmt.Println("=========")
	a := 0
	for ; a < 10; a++ {
		fmt.Println(a)
	}
	// 3. 省略初始语句和结束语句
	fmt.Println("=========")
	b := 10
	for b > 0 {
		fmt.Println(b)
		b--
	}
}
