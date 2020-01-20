package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

// 结构体初始化
func main() {
	//1.键值对初始化
	p1 := person{
		name: "sunny",
		city: "杭州",
		age:  28,
	}
	fmt.Println(p1)
	//2.值的列表进行初始化,这个顺序要跟结构体保持一致且每个值都要赋值
	p2 := person{
		"sunny",
		"南京",
		28,
	}
	fmt.Println(p2)
}
