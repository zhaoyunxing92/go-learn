package main

import "fmt"

// 切片
func main() {
	//声明和初始化
	var g []int  // 声明切片
	h := []int{} // 声明并初始化

	if g == nil {
		fmt.Println("g is nil") //g is nil
	}
	fmt.Println(g, cap(g), len(g)) //[] 0 0

	if h == nil {
		fmt.Println("h is nil")
	}
	fmt.Println(h, cap(h), len(h)) // [] 0 0

	// 1.基于数组的切片
	a := [5]int{10, 11, 12, 13, 14}
	b := a[1:3]           //a数组的1~3组成数组b
	fmt.Println(a)        //[10 11 12 13 14]
	fmt.Println(b)        //[11 12]
	fmt.Printf("%T\n", b) // []int
	// 2.切片基础上再次切片
	c := b[:]
	fmt.Println(c) //[11 12]
	// 通过make创建
	d := make([]int, 5, 10)
	fmt.Println(d)                 //[0 0 0 0 0]
	fmt.Println("数组d的长度:", len(d)) //数组d的长度: 5
	fmt.Println("数组d的容量:", cap(d)) //数组d的容量: 10

	// 切片的赋值和拷贝
	e := make([]int, 3, 3)
	f := e
	f[0] = 100     //修改f切片会导致e也改变
	fmt.Println(e) //[100 0 0]
	fmt.Println(f) //[100 0 0]
	//append
	var j []int //声明切片但是没有申请内存
	j = append(j, 10, 20, 30, 40)
	fmt.Println(j) // [10 20 30 40]

	// 切片的cope
	m := []int{1, 2, 3, 4, 5}
	k := make([]int, 5, 5)
	copy(k, m)
	k[0] = 100
	fmt.Println(m) //[1 2 3 4 5]
	fmt.Println(k) //[100 2 3 4 5]

}
