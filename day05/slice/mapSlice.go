package main

import "fmt"

// map的切片
func main() {
	// 切片初始化
	mapSlice := make([]map[string]int, 3, 8)

	// map初始化
	mapSlice[0] = make(map[string]int, 8)

	// 赋值
	mapSlice[0]["sunny"] = 100

	fmt.Println(mapSlice) //[map[sunny:100] map[] map[]]
}
