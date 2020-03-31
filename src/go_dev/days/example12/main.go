package main

import (
	"fmt"
	"sort"
)

func testSort() {
	// var arr []int = []int{2, 3, 1, 5, 2, 4}
	var arr = []int{2, 3, 1, 5, 2, 4}
	sort.Ints(arr[:])
	fmt.Println(arr)
}

func testString() {
	var arr = []string{"a", "df", "rc", "fW"}
	sort.Strings(arr[:])
	fmt.Println(arr)
}

func modify(a map[string]map[string]string) {
	val, ok := a["zhangsan"]
	if ok {
		val["password"] = "123456"
		val["nickname"] = "ahua"
	} else {
		a["zhangsan"] = make(map[string]string)
		a["zhangsan"]["password"] = "123456"
		a["zhangsan"]["nickname"] = "ahua"
	}
	return
}

func main() {
	var a map[string]string = map[string]string{
		"b": "bcd",
	}
	var b map[string]int
	var c map[int]int
	var d map[string]map[string]string
	fmt.Println(a, b, c, d)

	// a = make(map[string]string, 5)
	a["a"] = "abc"
	v, ok := a["a1"]
	fmt.Println(v, ok)
	fmt.Println(a, b, c, d)

	testSort()
	testString()
}
