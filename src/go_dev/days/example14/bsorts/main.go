package main

import "fmt"

func bsorts(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] > a[j-1] {

				// a[i], a[j] = a[j], a[i]
				a[j-1], a[j] = a[j], a[j-1]
				// a[j], a[j-1] = a[j-1], a[j]
				// temp := a[j-1]
				// a[j-1] = a[j]
				// a[j] = temp
			}
		}
	}
}

func main() {
	// var a []int = []int{4, 2, 5, 1, 6, 3}
	a := [...]int{4, 2, 5, 1, 6, 3}
	fmt.Println(a)
	bsorts(a[:])
	fmt.Println(a)
}
