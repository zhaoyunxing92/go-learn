# 匿名函数和闭包

## 匿名函数定义

匿名函数就是没有函数名称，匿名函数多用于实现回调函数和闭包

## 语法

```go
func(参数)(返回值){
  函数体
}
```
## 实战部分

### 赋值给一个变量

```go
sayHello := func() {
    fmt.Println("匿名函数")
}
sayHello()
```

### 立即执行的匿名函数

```go
func() {
    fmt.Println("立即执行")
}()
```

## 闭包

* 闭包= 函数+ 外层变量的引用
* 返回值是一个函数

### 实战部分

#### 判断字符串结尾,如果是则返回，如果不是则拼接结尾后返回

```go
func main() {
    r := makeSuffixFunc(".txt")
    ret := r("zhaoyunxing")
    fmt.Println(ret)
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(str string) string {
        //判断字符串的结尾是否符合要求
		if !strings.HasSuffix(str, suffix) {
			return str + suffix
		}
		return str
	}
}
```

#### 变量复用

```go
func main() {
 	x, y := calc(100)
 	ret1 := x(200) //300
 	ret2 := y(100) //200
 	fmt.Println(ret1, ret2)
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
```


