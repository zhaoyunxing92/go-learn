package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum    int64
	lock   sync.Mutex
	rwLock sync.RWMutex
	wg     sync.WaitGroup
)

func read() {
	time.Sleep(time.Microsecond)
	wg.Done()
}

func write() {
	sum += 1
	time.Sleep(time.Microsecond * 10)
	wg.Done()
}

// 互斥锁写
func mutexWrite() {
	lock.Lock()
	write()
	lock.Unlock()
}

func mutexRead() {
	lock.Lock()
	read()
	lock.Unlock()
}

func rwMutexRead() {
	rwLock.RLock()
	read()
	rwLock.RUnlock()
}

func rwMutexWrite() {
	rwLock.Lock()
	write()
	rwLock.Unlock()
}

//读写锁
func main() {
	start := time.Now()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go mutexRead()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mutexWrite()
	}
	wg.Wait()
	fmt.Printf("sum:%d,mutex time:%s \n", sum, time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go rwMutexRead()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go rwMutexWrite()
	}
	wg.Wait()
	fmt.Printf("sum:%d,rw mutex time:%s \n", sum, time.Now().Sub(start))
}
