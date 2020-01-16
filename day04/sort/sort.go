package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [...]int{10, 19, 28, 13, 14}
	sort.Ints(a[:])
	fmt.Print(a) //[10 13 14 19 28]
}
