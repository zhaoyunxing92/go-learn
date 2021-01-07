package main

import "fmt"

// map(映射)
func main() {
	//初始化容量为8的map集合
	a := make(map[string]int, 8)
	a["sunny"] = 100
	a["zhaoyunxing"] = 200

	// 输出
	fmt.Printf("a:%#v\n", a)   //a:map[string]int{"sunny":100, "zhaoyunxing":200}
	fmt.Printf("type:%T\n", a) //type:map[string]int

	// 声明并初始化
	b := map[int]bool{
		1: true,
		2: false,
		3: true,
	}
	fmt.Printf("a:%#v\n", b)    //a:map[int]bool{1:true, 2:false, 3:true}
	fmt.Printf("type:%T\n", b)  //type:map[int]bool
	fmt.Println("len:", len(b)) //len: 3

	// 判断某个键是否存在
	score := make(map[string]int, 8)
	score["sunny"] = 95
	score["zhaoyunxing"] = 98
	score["lisi"] = 91

	// 判断是否存在
	val, ok := score["sunny"]
	if ok {
		fmt.Println("val:", val)
	}

	// for range 遍历
	for key := range score {
		fmt.Println(key)
	}

	for key, val := range score {
		fmt.Println(key, val)
	}

	// 删除指定key
	delete(score, "sunny")

}
