# 结构体

go语言提供自定义的数据类型，可以封装基本数据类型，这种数据类型叫结构体(`struct`),可以理解为java中的对象

## 语法

```go
type 类型名称 struct{
  字段名  字段类型
  字段名  字段类型
  ....
}
```

### 案例

#### 定义

```go
type person struct {
	name string
	age  int8
	city string
}
/**
字段类型如果相同的话可以写在一行
*/
type man struct {
	name, city string
	age        int8
}
```

#### 实例化

```go
p := person{} //var p person
p.name = "sunny"
p.city = "杭州"
p.age = 28
```

#### 匿名结构体

```go
var user struct {
    name string
    age  uint8
}
user.name = "sunny"
user.age = 28
```

#### 结构体指针

```go
var ps = new(person)
fmt.Printf("%T\n", ps) //*main.person
ps.name = "sunny"
ps.city = "杭州"
ps.age = 28
fmt.Printf("%#v\n", ps) //&main.person{name:"sunny", city:"杭州", age:28}
```

#### 取结构体地址实例化

```go
p1 := &person{}
p1.age = 28
p1.name = "sunny"
fmt.Printf("%#v\n", p1) //&main.person{name:"sunny", city:"", age:28}
```

## 结构体初始化

### 键值对初始化

```go
p1 := person{
    name: "sunny",
    city: "杭州",
    age:  28,
}
fmt.Println(p1)
```

### 值列表初始化

> 值列表初始化的顺序必须要跟结构体保持一致并且全部字段都要初始化

```go
p2 := person{
    "sunny",
    "南京",
    28,
}
fmt.Println(p2)
```

## 结构体的构造函数

go语言是没构造函数的，但是可以自己实现，由于结构体是值类型（内存开销比较大），所以构造函数返回的是结构体的指针类型这样比较节约内存