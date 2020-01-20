package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

/**
构造函数
*/
func newPerson(name, city string, age int8) *person {
	return &person{name: name, city: city, age: age}
}

// 构造函数
func main() {
	ps := newPerson("sunny", "杭州", 28)
	fmt.Printf("type:%T value:%#v\n", ps, ps) //type:*main.person value:&main.person{name:"sunny", city:"杭州", age:28}
}
