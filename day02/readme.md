# 运算符和流程控制

## 运算

### 算术运算

```go
a := 20
b := 10

fmt.Println(a + b) //30
fmt.Println(a - b) //10
fmt.Println(a * b) // 200
fmt.Println(a / b) // 2
fmt.Println(a % b) // 0
```

### 关系运算

```go
fmt.Println(20 > 10)  //true
fmt.Println(20 != 10) //true
fmt.Println(20 <= 10) //false
```

### 逻辑运算

```go
fmt.Println(20 > 10 && 10 == 10) // true
fmt.Println(!(1 > 5))            //true
fmt.Println((1 > 5) || (5 > 2))  //true
```

### 位运算

```go
g := 1              //001
h := 5              //101
fmt.Println(g & h)  // 001 => 1
fmt.Println(g | h)  // 101 => 5
fmt.Println(g ^ h)  // 100 => 4
fmt.Println(g ^ h)  // 100 => 4
fmt.Println(1 << 2) // 100 => 4
fmt.Println(4 >> 2) // 001 => 1
```

### 赋值运算

```go
var e int
e = 5
e += 8         //e=e+8
fmt.Println(e) //13
```

## 流程控制

```go
package main
import "fmt"
// if 流程控制
func main() {
	//1.基础写法
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//2.if特殊写法 s 变量只在当前语句有效
	if s := 65; s >= 90 {
		fmt.Println("A")
	} else if s >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

## for循环

### 基础写法

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### 忽略初始语句

```go
a := 0
for ; a < 10; a++ {
    fmt.Println(a)
}
```

###  省略初始语句和结束语句

```go
b := 10
for b > 0 {
    fmt.Println(b)
    b--
}
```