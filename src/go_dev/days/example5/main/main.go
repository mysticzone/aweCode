package main

import (
	"fmt"
)

var a int

//a := 100
//var a int = 100

func test(){
	var b int8 = 200
	var c int32

	c = int32(b)
	fmt.Println(b)
	fmt.Println(c)
}




func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}

func swap1(a int, b int) (int,int) {
	return b, a

}

func main(){
	var a int = 10
	var b int = 20

	//swap(&a, &b)
	//fmt.Println("a: ", a)
	//fmt.Println("b: ", b)

	a, b = swap1(a, b)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
