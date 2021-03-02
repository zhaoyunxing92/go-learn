package main

/**
定义一个人
*/
type person struct {
	name string
	city string
	age  uint8
}

/**
字段类型如果相同的话可以写在一行
*/
type man struct {
	name, city string
	age        int8
}

// 结构体
func main() {

	p := person{} //var p person
	p.name = "sunny"
	p.city = "杭州"
	p.age = 28

	// 匿名结构体
	var user struct {
		name string
		age  uint8
	}
	user.name = "sunny"
	user.age = 28
}
