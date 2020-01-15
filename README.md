# go-learn

## go 一般结构

 * 只有`package`名称为`main`的包可以包含`main`函数
 * 一个可执行程序`有且仅有一个main包`
 * 常量定义使用`const`
 * 全局变量声明使用`var`
 * 一般类型声明使用`type` 
 * 结构声明`struct`
 * 接口声明`interface`
 * 未使用的包必须删除

 ## go 可见性规则

 * go 通过首字母大小写决定常量、变量、类型、结构、函数是否可以被外部访问，大写的可以被外部访问

 ## go 变量组声明方式

  ```go
// 常量组定义
const(
    name="sunny"
    name2="sunny2"
)
//全局变量声明和赋值
var(
    name="sunny"
    age=27
)
//一般类型声明
type(
    type1 string
    type2 int
)
  ```

## go 基本类型

* 布尔值:bool

  - 长度：１字节
  - 取值：true/false
  - 默认值：false
  
* 整型：int/uint

  - 根据平台可能是32或64位
  
* 8位整型:int8/uint8

  - 长度: 1字节
  - 取值: -128&sim;127 / 0&sim;255
  
* 字节型:byte(uint8别名)

* 16位整型：int16/uint16

  - 长度：2字节
  - 取值：-32768&sim;32767/0&sim;65535

* 32位整数:int32/uint32
  - 长度：4字节
  - 取值：-2^32/2&sim;2^32/2-1/0&sim;2^32-1
  
* 64位整数:int64/uint64 
  - 长度：8字节
  - 取值：-2^64/2&sim;2^64/2-1/0&sim;2^64-1

* 浮点型:float32/float64
  - 长度：4/8字节
  - 取值：小数位７/１５位

* 复数: complex64/complex128
  - 长度：8/16字节

* 其它值类型

  - array 、struct、string
  
* 引用类型

  - slice、map、chan

* 函数类型

  - func

* 接口类型

  - inteface

* 足够保持指针的32位或64位的整数型：uintptr