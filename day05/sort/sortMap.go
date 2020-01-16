package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// map key有序遍历
func main() {
	// 初始化数据
	score := make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i) //保留两位小数
		value := rand.Intn(100)          //0~99的随机数
		score[key] = value
	}

	// 提取全部的key存放到切片中
	keys := make([]string, 0, 100)
	for key := range score {
		keys = append(keys, key)
	}
	//切片排序
	sort.Strings(keys)

	//遍历切片
	for _, key := range keys {
		fmt.Println(key, score[key])
	}
}
