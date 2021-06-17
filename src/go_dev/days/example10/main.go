package main

import "fmt"

func fab(n int) {
	var a []int
	a = make([]int, n)

	a[0] = 1
	a[1] = 1

	var age0 [5]int = [5]int{1, 2, 3, 4, 5}
	var age1 = [5]int{1, 2, 3, 4, 5}
	var age2 = [...]int{1, 2, 3, 4, 5}
	var str = [5]string{3: "hello world", 4: "tom"}
	fmt.Println(age0, age1, age2, str)

}

func modify(a *[6]int) {
	fmt.Println(a)
	a[0] = 101
	fmt.Println(a)

}

func main() {
	var a [6]int
	a[0] = 100
	fmt.Println(&a)

	modify(&a)
	fmt.Println(a)

	fab(10)
}
