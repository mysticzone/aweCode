package main

import "fmt"

func Ptr1() {
	var x int
	var x_ptr *int

	fmt.Println("x_ptr: ", x_ptr)

	x = 10
	x_ptr = &x
	fmt.Println("x: ", x)
	fmt.Println("x_ptr: ", x_ptr)
	fmt.Println("*x_ptr: ", *x_ptr)
}

func change(x *int) {
	*x = 200
}

func PtrChange() {
	var x int = 100
	fmt.Println(x)

	change(&x)
	fmt.Println(x)

}

func SetValue(x *int) {
	*x = 100
}

func PtrSet() {
	x_ptr := new(int)
	SetValue(x_ptr)

	fmt.Println("x_ptr: ", *x_ptr)
	fmt.Println("x_ptr: ", &x_ptr)
	fmt.Println("x_ptr: ", x_ptr)

}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func PtrSwap() {
	var a int = 100
	var b int = 200
	fmt.Println("a: ", a, "b: ", b)
	Swap(&a, &b)
	fmt.Println("a: ", a, "b: ", b)

}

func main() {
	Ptr1()
	PtrChange()
	fmt.Println("---------------------------")
	PtrSet()
	fmt.Println("---------------------------")
	PtrSwap()
}
