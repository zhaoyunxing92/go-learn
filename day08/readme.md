# panic and recover

* panic可以理解为是java中的throw exception会中断程序的执行。

* recover可以理解为是java中的try catch

## recover使用须知

* recover必须在defer函数中使用

* 在正常函数执行过程中，调用recover函数没有任何作用，返回nil

* 如果函数发生了panic，那么recover会捕获这个panic的值，并且让程序正常执行下去，不会crash

## 案例

```go
func a() {
	fmt.Println("func a")
}

func b() {
	defer func() { // defer 必须先定义，否定无法捕获panic异常
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}
```