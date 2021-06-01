package main

import (
	"fmt"
	"time"
)

// 定时器
func main() {
	var ch chan int
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		ch <- 1
	}()
	<-ch
}
