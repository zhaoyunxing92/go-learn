# map

go中的map是一种无序的`key-value`数据结构，go的map是引用类型，必须初始化才能使用

## 语法

```go
map[keyType]valueType
```
* keyType: 表示键类型
* valueType: 值类型

map类型的变量默认初始值为nil,需要make()函数来分配内存

```go
make(map[keyType]valueType,[cap])
```

### 实战

#### 初始化容量为8的map集合

```go
a := make(map[string]int, 8)
a["sunny"] = 100
a["zhaoyunxing"] = 200

// 输出
fmt.Printf("a:%#v\n", a)   //a:map[string]int{"sunny":100, "zhaoyunxing":200}
fmt.Printf("type:%T\n", a) //type:map[string]int
```

#### 声明并初始化

```go
b := map[int]bool{
    1: true,
    2: false,
    3: true,
}
fmt.Printf("a:%#v\n", b)    //a:map[int]bool{1:true, 2:false, 3:true}
fmt.Printf("type:%T\n", b)  //type:map[int]bool
fmt.Println("len:", len(b)) //len: 3
```

#### 判断某个键是否存在

```go
score := make(map[string]int, 8)
score["sunny"] = 95
score["zhaoyunxing"] = 98
score["lisi"] = 91

// 判断是否存在
val, ok := score["sunny"]
if ok {
    fmt.Println("val:", val)
}
```

#### map遍历

```go
for key := range score {
    fmt.Println(key)
}

for key, val := range score {
    fmt.Println(key,val)
}
```

#### 删除指定key

```go
delete(score,"sunny")
```

#### 排序key后遍历集合

```go
// map key有序遍历
func main() {
	// 初始化数据
	score := make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i) //保留两位小数
		value := rand.Intn(100)          //0~99的随机数
		score[key] = value
	}

	//提取全部的key存放到切片中
	keys := make([]string, 0, 100)
	for key := range score {
		keys = append(keys, key)
	}
	//切片排序
	sort.Strings(keys)

	//遍历切片
	for _, key := range keys {
		fmt.Println(key, score[key])
	}
}
```

#### 统计字符串中单词的出现次数

```go
//统计how do you do字符串中单词出现的频率
func main() {
	str := "how do you do"
	//字符串空格分割
	words := strings.Split(str, " ")
	// 定义一个map
	wordsCount := make(map[string]int, 8)
	// 遍历分割后的单词
	for _, word := range words {
		// 判断是否存在
		val, ok := wordsCount[word]
		if ok {
			wordsCount[word] = val + 1
		} else {
			wordsCount[word] = 1
		}
	}
	// 遍历统计后结果
	for key, val := range wordsCount {
		fmt.Println(key, val)
	}
}
```

### 特殊的map

#### map的切片

```go
func main() {
    // 切片初始化
	mapSlice := make([]map[string]int, 3, 8)
	// map初始化
	mapSlice[0]=make(map[string]int,8)
	// 赋值
	mapSlice[0]["sunny"]=100

	fmt.Println(mapSlice) //[map[sunny:100] map[] map[]]
}
```

#### 值是切片的map

```go
//值是切片的map
func main() {
	// map初始化
	sliceMap := make(map[string][]int, 8)
	val, ok := sliceMap["sunny"]
	if ok {
		fmt.Println(val)
	} else {
		// 初始化切片
		sliceMap["sunny"] = make([]int, 3, 8)
		sliceMap["sunny"][0] = 100
		sliceMap["sunny"][2] = 100
	}
    // 遍历
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}
```