package main

import "fmt"

// 运算
func main() {
	// 1.算术运算
	a := 20
	b := 10

	fmt.Println(a + b) //30
	fmt.Println(a - b) //10
	fmt.Println(a * b) // 200
	fmt.Println(a / b) // 2
	fmt.Println(a % b) // 0

	// 2.关系运算符
	fmt.Println(20 > 10)  //true
	fmt.Println(20 != 10) //true
	fmt.Println(20 <= 10) //false

	//3.逻辑运算 && 、 || 、 !
	fmt.Println(20 > 10 && 10 == 10) // true
	fmt.Println(!(1 > 5))            //true
	fmt.Println((1 > 5) || (5 > 2))  //true

	// 4.位运算 &、| 、^ 、<< 、>>
	fmt.Println("位运算")
	g := 1              //001
	h := 5              //101
	fmt.Println(g & h)  // 001 => 1
	fmt.Println(g | h)  // 101 => 5
	fmt.Println(g ^ h)  // 100 => 4
	fmt.Println(g ^ h)  // 100 => 4
	fmt.Println(1 << 2) // 100 => 4
	fmt.Println(4 >> 2) // 001 => 1

	// 5.赋值运算
	fmt.Println("赋值运算")
	var e int
	e = 5
	e += 8         //e=e+8
	fmt.Println(e) //13
}
