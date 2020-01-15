package main

import "fmt"

// switch
func main() {
	// 1.基础用法
	score := 70
	switch score {
	case 100:
		fmt.Println("甲")
	case 70:
		fmt.Println("乙")
	case 60:
		fmt.Println("丙")
	default:
		fmt.Println("丁")
	}
	//2. case一次多个值
	a := 5
	switch a {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("未定义")
	}

	// case中使用表达式
	age := 30
	switch {
	case age > 18:
		fmt.Println("成年人")
	case age <= 18:
		fmt.Println("未成年人")
	}

}
