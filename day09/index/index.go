package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d,ptr:%p\n", a, &a)
	fmt.Println(b)  // 打印内存地址
	fmt.Println(*b) //获取内存地址存的值
	fmt.Println(&b) // 内存地址的地址

	name(a)
	fmt.Println(a)

	num:=50
	modify1(num)
	fmt.Println(num)
	modify2(&num)
	fmt.Println(num)
}

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 200
}

func name(a int) {
	a = 30

}
