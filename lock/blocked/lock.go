package main

import (
	"sync"
	"time"
)

var gbLock sync.Mutex

func main() {
	go func() {
		doubleRlock()
	}()
	for i := 0; i < 10; i++ {
		gbLock.Lock()
		gbLock.Unlock()
		time.Sleep(time.Millisecond)
	}

}
func doubleRlock() {
	gbLock.Lock()

	time.Sleep(time.Second)
	//rlock()
	gbLock.Unlock()
}

func rlock() {
	gbLock.Lock()
	gbLock.Unlock()
}
