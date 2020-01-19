package main

import "fmt"

func a() {
	fmt.Println("func a")
}

func b() {
	defer func() { // defer 必须先定义，否定无法捕获panic异常
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}
