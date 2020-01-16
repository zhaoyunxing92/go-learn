# 函数

函数是组织好的、可重复使用的、用于执行指定任务的代码块，go语言中支持函数、匿名函数、闭包

## 语法

```go
func 函数名(参数)(返回值){
  函数体
}
```

## 实战

### 无参无返回

```go
func sayHello() {
	fmt.Println("hello go")
}
```

### 带参数

```go
func sayHello2(name string) {
	fmt.Println("hello", name)
}
```

### 单个返回值

```go
func intSum(a int, b int) int {
	return a + b
}
```

### 多个返回值

```go
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
```

### 可变参数

在函数体中可变参数类型是切片,并且可变参数必须在最后

```go
func intSum3(a ...int) int {
	sum := 0
	for _, arg := range a {
		sum += arg
	}
	return sum
}
```

### 参数类型简写

如果参数a和参数b的类型相同则只用最后写一个

```go
func intSum4(a, b int) int {
	return a + b
}
```

### 返回值简写

go会自动返回变量名称是sum的变量

```go
func intSum2(a int, b int) (sum int) {
	sum = a + b
	return
}
```

### 参数是一个函数

```go
package main

import "fmt"
/**
求和
*/
func add(x, y int) int {
	return x + y
}
/**
差值
*/
func sub(x, y int) int {
	return x - y
}
/**
函数作为参数
*/
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	// 第三个参数是函数
	r1 := calc(100, 200, add)
	r2 := calc(100, 200, sub)
	fmt.Println(r1)
	fmt.Println(r2)
}
```

### 函数访问变量规则

现在函数内部找，如果找到了则用，没有就找全局变量

```go
package main
import "fmt"
var num = 10
func testGlobal() {
	num = 20
	fmt.Println("全局变量：", num)
}
func main() {
	testGlobal() // 20
}
```

## defer

defer是指延迟执行，在一个函数即将执行完时，才倒序执行defer定义的函数.

```go
func main() {
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")
	//执行结果
	//start...
	//end...
	//3
	//2
	//1
}
```

### defer使用场景

由于defer的延迟调用特性，很方便的用于资源释放、文件关闭、解锁、记录时间等场景
