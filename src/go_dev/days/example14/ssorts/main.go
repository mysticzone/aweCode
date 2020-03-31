package main

import "fmt"

func ssorts(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

}

func main() {
	// var a []int = []int{4, 2, 5, 1, 6, 3}
	a := [...]int{4, 2, 5, 1, 6, 3}
	fmt.Println(a)
	ssorts(a[:])
	fmt.Println(a)
}
