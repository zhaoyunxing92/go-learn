package main

import "fmt"

/*
常量
* 使用cost关键字
* 必须指定值
* 如果一组常理只定义了一个，下面的默认跟第一个相同
* iota 计数器
*/
const pi = 3.14

const (
	a = 10 //10
	b      //10
	c      //10
)

const (
	n1 = iota // 0
	n2        // 1
	n3        // 2
	n4        // 3
)

// 跳过某个值
const (
	a1 = iota //0
	a2        //1
	_         // 跳过
	a4        //3
)

const (
	b1 = iota //0
	b2 = 100  //100
	b3 = iota //2
	b4        //3

)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //1<< 10  1024
	MB = 1 << (10 * iota) //1<< 20  1048576
	GB = 1 << (10 * iota) //1<< 30  1073741824
	TB = 1 << (10 * iota) //1<< 40  1099511627776
)

// 多个iota定义在一行

const (
	e, f = iota + 1, iota + 2 //1 2
	g, h                      //2 3
	w, q                      //3 4
)

func main() {

	fmt.Println(a, b, c) //10 10 10

	fmt.Println(n1, n2, n3, n4) //0 1 2 3

	fmt.Println(a1, a2, a4) //0 1 3

	fmt.Println(b1, b2, b3, b4) //0 1 3

	fmt.Println(KB, MB, GB, TB) //1024 1048576 1073741824 1099511627776

	fmt.Println(e, f, g, h, w, q) //1 2 2 3 3 4
}
