package main

import "fmt"
import "time"

func add(a, b int) (c int) {
	c = a + b
	return
}

func args(a, b int, arg ...int) {
	fmt.Println(a, b)
	for i := 0; i < len(arg); i++ {
		fmt.Println(arg[i])
	}
}

func test(a int, b int) {
	result := func(a int, b int) int {
		return a + b
	}
	fmt.Println(result(a, b))
}

func main() {

	test(2, 2)
	args(0, 0, 1, 2, 3, 4)

	var i int = 0
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println(1)

	fmt.Println("--------------")
	i = 10
	fmt.Println(i)
	fmt.Println()

	fmt.Println()
	time.Now()

}
