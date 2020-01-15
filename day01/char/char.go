package main

import "fmt"

func main() {
	/* 字符
	 * byte uint8的别名 ASCII码
	 * rune int32的别名
	 */
	var a1 byte = 'a' //97
	var a2 rune = 'a' //97
	fmt.Println(a1, a2)
	// 查看类型
	fmt.Printf("a1:%T, a2:%T", a1, a2) //a1:uint8,a2:int32
}
