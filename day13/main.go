package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			for c := 0; c < 1000; c++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					fmt.Println(a, b, c)
				}
			}
		}
	}
	end := time.Now().Sub(start)
	fmt.Println(end)
}
