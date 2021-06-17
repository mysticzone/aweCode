package main

import (
	"fmt"
	"strings"
)

func main() {
	var f = Adder()
	fmt.Println("func: ", f)
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(300))

	func1 := makeSuffixFunc(".bmg")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test1"))
	fmt.Println(func2("test2"))

	var a [10]int
	a[0] = 100
	for index, v := range a {
		fmt.Println(index, v)
	}
	fmt.Println("length: ", len(a))
	fmt.Println(a)
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func Adder() func(int) int {
	var x int
	ret := func(delta int) int {
		fmt.Println("> ", x)
		x += delta
		return x
	}
	fmt.Println(ret)
	// return ret
	return ret
}
