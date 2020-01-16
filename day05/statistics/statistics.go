package main

import (
	"fmt"
	"strings"
)

//统计how do you do字符串中单词出现的频率
func main() {
	str := "how do you do"
	//字符串空格分割
	words := strings.Split(str, " ")
	// 定义一个map
	wordsCount := make(map[string]int, 8)
	// 遍历分割后的单词
	for _, word := range words {
		// 判断是否存在
		val, ok := wordsCount[word]
		if ok {
			wordsCount[word] = val + 1
		} else {
			wordsCount[word] = 1
		}
	}
	// 遍历统计后结果
	for key, val := range wordsCount {
		fmt.Println(key, val)
	}
}
