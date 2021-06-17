package main

import "fmt"

func testString() {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	fmt.Println(cap(slice))

	slice = arr[2:4]
	slice[0] = 100
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println("-----")
	fmt.Println(&slice[0])
	fmt.Println(&arr[2])

}

type slice struct {
	prt *[]int
	len int
	cap int
}

func testSlice() {
	var a [4]int = [...]int{1, 2, 3, 4}
	s := a[:4]
	// s[3] = 100

	s = append(s, 100)
	s = append(s, 101)
	s[1] = 1000
	fmt.Println(s)
	fmt.Print(a)
}

func testCopy() {
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := make([]int, 10)
	fmt.Println(b)

	copy(b, a)
	fmt.Println(a)

	fmt.Println(b)
}

func main() {
	// testString()
	// fmt.Println("------------")
	// testSlice()
	fmt.Printf("------------\n")
	// testCopy()

	s := "hello"
	s1 := []rune(s)
	s1[0] = 'o'
	fmt.Println(s1)
	fmt.Println(s)

}
