# 类型别名和自定义类型

go语言中有一些基本数据类型如`string`、`int`、`bool`等数据类型，但是可以使用`type`关键字自定义类型

## 类型别名

> 类型别名是`go1.9`版本添加的功能

## 实战部分
```go
//1.自定义类型
type MyInt int

// 定义别名
type NewInt = int

//类型别名和自定义类型
func main() {

	var a MyInt
	fmt.Printf("type:%T value:%v\n", a, a) //type:main.MyInt value:0
	
	var b NewInt
	fmt.Printf("type:%T value:%v\n", b, b) //type:int value:0
}
```