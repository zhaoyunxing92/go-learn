package main

import "fmt"

// if 流程控制
func main() {
	//1.基础写法
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//2.if特殊写法 s 变量只在当前语句有效
	if s := 65; s >= 90 {
		fmt.Println("A")
	} else if s >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
