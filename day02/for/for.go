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
	//4.无限循环
	for {
		fmt.Println("1")
	}

	// 5.break跳出循环
	fmt.Println("=========")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}
	// 6.continue 继续下次循环
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
