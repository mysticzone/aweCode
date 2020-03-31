package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var lock sync.Mutex

func testMap() {
	var a map[int]int
	a = make(map[int]int, 10)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	a[10] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	fmt.Println(a)
}

func main() {
	testMap()
}
