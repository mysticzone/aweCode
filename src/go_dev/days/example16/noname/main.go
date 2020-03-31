package main

import (
	"fmt"
	"time"
)

type Cart struct {
	name string
	age  int
}

type Train struct {
	Cart
	int
	start time.Time
}

func main() {
	var t Train
	t.name = "huoche"
	t.int = 100
	t.age = 99

	fmt.Println(t)
}
