# 切片
切片（slice）是建立在数组之上的更方便，更灵活，更强大的数据结构。切片并不存储任何元素而只是对现有数组的引用。

## 切片的声明

### 声明和初始化

```go
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
```

### 基于数组

```go
a := [5]int{10, 11, 12, 13, 14}
b := a[1:3]           //a数组的1~3组成数组b
fmt.Println(a)        //[10 11 12 13 14]
fmt.Println(b)        //[11 12]
fmt.Printf("%T\n", b) // []int
```

### 在切片基础上再次切片

```go
c := b[:]
fmt.Println(c) //[11 12]
```

### 使用make函数

```go
d := make([]int, 5, 10)
fmt.Println(d)                 //[0 0 0 0 0]
fmt.Println("数组d的长度:", len(d)) //数组d的长度: 5
fmt.Println("数组d的容量:", cap(d)) //数组d的容量: 10
```
## 切片的赋值和拷贝

```go
e := make([]int, 3, 3)
f := e
f[0] = 100     //修改f切片会导致e也改变
fmt.Println(e) //[100 0 0]
fmt.Println(f) //[100 0 0]
```

## 切片的遍历参考数组的遍历

## 切片的append

当append超过元素的容量时，会进行扩容，所以必须要接收append后的值

```go
var j []int //声明切片但是没有申请内存
j = append(j, 10, 20, 30, 40) // j变量接收append后的切片
fmt.Println(j) // [10 20 30 40]
```
## 切片的copy

copy后的数组互不影响

```go
m := []int{1, 2, 3, 4, 5}
k := make([]int, 5, 5)
copy(k, m)
fmt.Println(m) //[1 2 3 4 5]
fmt.Println(k) //[100 2 3 4 5]
```

## 切片排序

```go
a := [...]int{10, 19, 28, 13, 14}
sort.Ints(a[:])
fmt.Print(a) //[10 13 14 19 28]
```