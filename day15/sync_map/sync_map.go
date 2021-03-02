package main

import (
	"fmt"
	"sync"
)

var (
	m  sync.Map
	wg sync.WaitGroup
)

// 并发map，可以多个协程同时调用
func main() {

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(idx int) {
			m.Store(idx, idx*100)
			value, _ := m.Load(idx)
			fmt.Printf("key:%v,value:%v \n", idx, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
