package main

import "fmt"

func isorts(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				break
			}
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
}

func main() {
	// var a []int = []int{4, 2, 5, 1, 6, 3}
	a := [...]int{4, 2, 5, 1, 6, 3}
	fmt.Println(a)
	isorts(a[:])
	fmt.Println(a)
}
