package main

import (
	"fmt"
	"strings"
)

// 字符串常见操作
func main() {
	// 字符串长度
	str := "hello"
	fmt.Println(len(str)) // 5

	str1 := "hello赵云兴"
	fmt.Println(len(str1)) // 14

	str2 := "赵云兴"
	fmt.Println(len(str2)) // 9 说明一个汉字占3个长度

	// 字符串拼接
	a := "hello"
	b := "world"
	c := fmt.Sprintf("%s %s", a, b)
	fmt.Println(c) // hello world

	// 字符串分割
	d := "hello&go"
	fmt.Println(strings.Split(d, "&"))        //[hello go]
	fmt.Printf("%T\n", strings.Split(d, ",")) //查看分割后的类型：[]string

	// 判断是否包含
	fmt.Println(strings.Contains(d, "go")) //true

	// 获取字符串位置 第一次出现的位置，不存在返回-1
	fmt.Println(strings.Index(d, "go"))    //6
	fmt.Println(strings.IndexAny(d, "o"))  // 4
	fmt.Println(strings.LastIndex(d, "o")) // 7

	// join 拼接
	e := []string{"hello", "world"}
	fmt.Println(strings.Join(e, " ")) //hello world

}
