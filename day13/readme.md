# interface 接口

go中的接口是一种类型，一种抽象的类型

## 代码
```go
//说话的类型
type sayer interface {
	say()
}

type cat struct {
}

type dog struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

// go中的接口是一种类型，一种抽象的类型
func main() {

}
```