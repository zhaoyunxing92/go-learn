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

### 无限循环

```go
for {
    fmt.Println("1")
}
```

### break 跳出循环

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
    if i == 3 {
        break
    }
}
```

### continue跳出本次循环

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue
    }
    fmt.Println(i)
}
```

## switch

### 基础用法

```go
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
```

### case多个值

```go
a := 5
switch a {
case 1, 3, 5, 7, 9:
    fmt.Println("奇数")
case 2, 4, 6, 8:
    fmt.Println("偶数")
default:
    fmt.Println("未定义")
}
```

### case只用表达式

```go
age := 30
switch {
case age > 18:
    fmt.Println("成年人")
case age <= 18:
    fmt.Println("未成年人")
}
```