package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

//结构体指针
func main() {
	// 采用new实例化
	var ps = new(person)
	fmt.Printf("%T\n", ps) //*main.person
	ps.name = "sunny"
	ps.city = "杭州"
	ps.age = 28
	fmt.Printf("%#v\n", ps) //&main.person{name:"sunny", city:"杭州", age:28}

	// 取结构体地址实例化
	p1 := &person{}
	p1.age = 28
	p1.name = "sunny"
	fmt.Printf("%#v\n", p1) //&main.person{name:"sunny", city:"", age:28}

}
