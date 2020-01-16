package main

import "fmt"

//函数
func main() {

	sayHello()

	sayHello2("go")

	fmt.Println(intSum(5, 3))
	fmt.Println(intSum2(5, 3))
	fmt.Println(intSum3(1, 2, 3, 4, 5))

	a, b := calc(100, 200)
	fmt.Println(a, b) //300 -100
}

// 无参无返回
func sayHello() {
	fmt.Println("hello go")
}

// 带参数
func sayHello2(name string) {
	fmt.Println("hello", name)
}

//单个返回值的
func intSum(a int, b int) int {
	return a + b
}

//返回值简写 go会自动返回变量名称是sum的变量
func intSum2(a int, b int) (sum int) {
	sum = a + b
	return
}

// 可变参数,在函数体中是切片,并且可变参数必须在最后
func intSum3(a ...int) int {
	sum := 0
	for _, arg := range a {
		sum += arg
	}
	return sum
}

//类型简写,如果a和b的类型相同则只用最后写一个
func intSum4(a, b int) int {
	return a + b
}

// 多返回值函数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
