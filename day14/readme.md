# reflect 反射

```go
package main

import (
	"fmt"
	"reflect"
)

// 反射
func reflectType(x interface{}) {

	obj := reflect.TypeOf(x)
	fmt.Printf("type:%v,kind:%v\n", obj.Name(), obj.Kind())
}

type Cat struct{}

type Dog struct{}

func main() {

	var a float32 = 3.1415926

	reflectType(a) // type:float32,kind:float32

	var b int8 = 10

	reflectType(b) // type:int8,kind:int8

	var cat Cat

	reflectType(cat) //type:Cat,kind:struct

	var dog Dog

	reflectType(dog) //type:Dog,kind:struct

	// 切片、Map、指针、数组他们的.Name()都是空
	var slice []int

	reflectType(slice) //type:,kind:slice

	var str []string

	reflectType(str) //type:,kind:slice

}

```

## 类型（type） & 种类（kind）

> 类型（type）是自定义的类型，而种类（kind）是底层数据类型

> go语言中的种类

```go
type Kind uint

const (
	Invalid Kind = iota  //无效
	Bool        // 布尔
	Int         // 有符号整型
	Int8        // 有符号8位整型 -128 ~ 127
	Int16       // 有符号16位整型 -32768 ~ 32767
	Int32       // 有符号32位整型 -2147483648 ~ 2147483647
	Int64       // 有符号64位整型 -9223372036854775808 ~ 9223372036854775807
	Uint        // 无符号整型
	Uint8       // 无符号8位整型 0 ~ 255
	Uint16      // 无符号16位整型 0 ~ 65535
	Uint32      // 无符号32位整型 0 ~ 4294967295
	Uint64      // 无符号64位整型 0 ~ 18446744073709551615
	Uintptr     // 指针
	Float32     // 单精度浮点数
	Float64     // 双精度浮点数     
	Complex64   // 64位复数
	Complex128  // 128位复数
	Array       // 数组
	Chan        // 通道
	Func        // 函数
	Interface   // 接口
	Map         // 映射
	Ptr         // 指针
	Slice       // 切片
	String      // 字符串
	Struct      // 结构体
	UnsafePointer //底层指针
)
```

## ValueOf()

> `reflect.ValueOf()`返回的是`reflect.Value`类型，其中包含了值信息

|          方法           |                             说明                             |
| :---------------------: | :----------------------------------------------------------: |
| Interface() interface{} |   将值以interface{}类型返回，可以通过类型断言转为指定类型    |
|       Int() int64       |           将值以int类型返回，所有有符号整型均可以            |
|      Uint() uint64      |           将值以uint类型返回，所有无符号整型均可以           |
|     Float() float64     | 将值以双精度（float64）类型返回，所有浮点（float32、float64） |
|       Bool() bool       |                         bool类型返回                         |
|     Bytes() []bytes     |                   字节数组[]bytes类型返回                    |
|     String() string     |                        字符串类型返回                        |

> 代码

```go
// 反射取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v)

	//获取值的类型
	k := v.Kind()
	switch k {

	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v,%T\n", ret, ret) //10,int32
	}
}
```

## Elem

> Elem()根据指针取值,go中不能直接set值必须指针传递才可以修改

```go
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v) //0xc00000a0c8,reflect.Value

	//Elem()根据指针取值
	k := v.Elem().Kind()
	switch k {

	case reflect.Int32:
		v.Elem().SetInt(50)
	}
}
```

## IsNil() & IsValid()

> `IsNil()`一般判断指针是否为空,`IsValid()`一般判断返回值是否有效

```go
func main() {

	var a *int

	fmt.Println("var a *int is nil", reflect.ValueOf(a).IsNil()) //true
	fmt.Println("var a *int is valid", reflect.ValueOf(a).IsValid()) //true

    // 实例化一个匿名结构体
    b:= struct {}{}
    //反射获取其中字段
    fmt.Println("是否存在name字段",reflect.ValueOf(b).FieldByName("name").IsValid())
    fmt.Println("是否存在name方法",reflect.ValueOf(b).MethodByName("name").IsValid())

    //map
    c:= map[string]int{}
    fmt.Println("map中是否存sunny在键：",reflect.ValueOf(c).MapIndex(reflect.ValueOf("sunny")).IsValid())
}
```
