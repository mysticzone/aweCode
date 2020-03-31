package main

import "fmt"

func qsort(a []int, left, right int) {
	if left >= right {
		return
	}

	val := a[left]
	k := left

	for i := left + 1; i < right; i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}
	a[k] = val
	qsort(a, left, k-1)
	qsort(a, k+1, right)
}

func main() {
	// var a []int = []int{4, 2, 5, 1, 6, 3}
	a := [...]int{4, 2, 5, 1, 6, 3}
	fmt.Println(a)
	qsort(a[:], 0, len(a))
	fmt.Println(a)
}
