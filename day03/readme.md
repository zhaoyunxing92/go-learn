# 数组

## 定义

### 基础用法

```go
var city = [4] string{"北京", "上海", "广州", "深圳"}
fmt.Println(city) //[北京 上海 广州 深圳]
```

### 编译器推导数组长度

```go
var boolArray = [...] bool{true, false, true, false}
fmt.Println(boolArray) //[true false true false]
```

### 使用索引值初始化

```go
var langArray = [...]string{1: "golang", 8: "java", 5: "python"}
fmt.Println(langArray)
fmt.Printf("%T\n", langArray) //[9]string 长度为9的数组
```

## 遍历

### 使用for

```go
for i := 0; i < len(city); i++ {
    fmt.Println(city[i])
}
```

### 使用 for range

```go
// 2.for range
for index, name := range city {
    fmt.Println(index, name)
}
// 3. for range 简写
for name := range city {
    fmt.Println(name)
}
```

## 二维数组

```go
cityName := [3][2]string{
    {"北京", "西安"},
    {"上海", "杭州"},
    {"广州", "深圳"},
}
fmt.Println(cityName) // [[北京 西安] [上海 杭州] [广州 深圳]]
```