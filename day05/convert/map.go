package main

import "fmt"

type Person struct {
	Address map[string]Address
}

type Address struct {
	Name string
	Desc string
}

func main() {

	p := &Person{}
	a := make(map[string]Address, 2)
	a["a"] = Address{Name: "杭州", Desc: "浙江杭州"}
	a["b"] = Address{Name: "上海", Desc: "上海市"}

	p.Address = a

	b := make(map[string]*Address, 2)
	for k, v := range p.Address {
		c := &v
		fmt.Println("k=", k, "v=", c)
		b[k] = c
	}

	println(b)
}
