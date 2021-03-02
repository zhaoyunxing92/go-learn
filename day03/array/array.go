package main

import "fmt"

//数组相关
func main() {

	var a [3]int
	var b [4]int
	fmt.Println(a) //[0 0 0]
	fmt.Println(b) //[0 0 0 0]

	// 1.数组初始化 数组长度必须是一个固定的值
	var city = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(city) //[北京 上海 广州 深圳]

	//2.编译器推导数组长度
	var boolArray = [...]bool{true, false, true, false}
	fmt.Println(boolArray) //[true false true false]
	// 3.使用索引值初始化
	var langArray = [...]string{1: "golang", 8: "java", 5: "python"}
	fmt.Println(langArray)
	fmt.Printf("%T\n", langArray) //[9]string 长度为9的数组

	// 遍历
	// 1.for遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	// 2.for range
	for index, name := range city {
		fmt.Println(index, name)
	}
	// 3. for range 简写
	for name := range city {
		fmt.Println(name)
	}

	//二维数组
	cityName := [3][2]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"广州", "深圳"},
	}
	fmt.Println(cityName) // [[北京 西安] [上海 杭州] [广州 深圳]]

}
