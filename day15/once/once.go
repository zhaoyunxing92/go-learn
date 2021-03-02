package main

import (
	"sync"
)

var icons map[string]string

var wg sync.WaitGroup

var once sync.Once

func loadIcons() {
	icons = map[string]string{
		"left":  "left.png",
		"up":    "up.png",
		"right": "right.png",
		"down":  "down.png",
	}
}

// 只加载一次
func icon(name string) string {
	once.Do(loadIcons)
	return icons[name]
}

func main() {
	wg.Add(2)
	go icon("left")
	wg.Done()
}
