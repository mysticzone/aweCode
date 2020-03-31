package main

import (
	"go_dev/package_example/calc"
	"fmt"
)

func main(){
	sum := calc.Add(1, 2)
	sub := calc.Sub(5, 1)
	fmt.Println("Sum: ", sum)
	fmt.Println("Sub: ", sub)
}
