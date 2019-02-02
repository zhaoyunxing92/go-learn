// 当前程序的包名称，必须在开头
package main

// 导入其他包，可以使用()导入一批包
// import "fmt"
import (
	"fmt"
)

//　常量
const Str = "sunny"

// 全局变量
var test = "sunny"

//结构声明
type sunny struct {
}

// 接口声明
type myInterface interface {
}

//类型声明
type age int

func main() {
	fmt.Println("hello go")
	fmt.Println("hello go")
}
