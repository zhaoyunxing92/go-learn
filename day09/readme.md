# go语言中的指针

go语言中的函数传参都是值拷贝(也就是说函数中使用的外部变量是无法修改的),go中的指针使用`&`表示获取地址,`*`根据地址取值

## 指针地址和指针类型

每个变量都拥有一个地址，这个地址代表变量在内存中的位置。go语言使用`&`对变量进行`取地址`操作。go语言中的值类型（int、float、bool、string、array、struct）都用对应的指针类型，如:`*int`、`*string`

## 实战部分

### 指针操作

```go
a := 10
b := &a
fmt.Printf("a:%d,ptr:%p\n", a, &a) //a:10,ptr:0xc00005e068

fmt.Println(b) // 0xc00005e068 打印内存地址
fmt.Println(*b) // 10 获取内存地址存的值
fmt.Println(&b) // 0xc000088018 内存地址的地址
```

### 修改参数

```go
func main() {
num:=50
modify1(num) 
fmt.Println(num) // 50,值拷贝修改不管用
modify2(&num) 
fmt.Println(num) // 200 内存地址传递，参数被修改
}

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 200
}
```