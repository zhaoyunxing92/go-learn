package main

import "fmt"

//值是切片的map
func main() {
	// map初始化
	sliceMap := make(map[string][]int, 8)
	val, ok := sliceMap["sunny"]
	if ok {
		fmt.Println(val)
	} else {
		// 初始化切片
		sliceMap["sunny"] = make([]int, 3, 8)
		sliceMap["sunny"][0] = 100
		sliceMap["sunny"][2] = 100
	}
	// 遍历
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}
