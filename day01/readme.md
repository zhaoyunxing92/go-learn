# 变量、常量和字符串

## 变量

```go
package main
import "fmt"
/**
变量
* 标准声明
* 批量声明
* 变量类型推到和初始值
* 短变量声明（只能用在函数中）语法：:=
* 匿名变量（不占用命名空间，不会分配内存，可以重复使用，想忽略某个值的时候使用）语法：_
*/
func main() {
	//标准声明
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	// 批量声明
	var (
		a string
		b int
		c bool
		d int32
	)
	fmt.Println(a, b, c, d)
	// 变量类型推到和初始值
	var name1, age1 = "sunny", 28
	fmt.Println(name1, age1)

	// 短变量声明（只能用在函数中）
	e := 40
	fmt.Println(e)

	// 匿名变量
	x, _ := foo() //忽略第二个返回值
	_, y := foo() // 忽略第一个返回值

	fmt.Println(x, y)

}

func foo() (int, string) {
	return 10, "sunny"
}
```

## 常量

```go
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
```

## 字符串
```go
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
```

## char

```go
package main

import "fmt"

func main() {
	/* 字符
	 * byte uint8的别名 ASCII码
	 * rune int32的别名
	 */
	var a1 byte = 'a' //97
	var a2 rune = 'a' //97
	fmt.Println(a1, a2)
	// 查看类型
	fmt.Printf("a1:%T, a2:%T", a1, a2) //a1:uint8,a2:int32
}
```