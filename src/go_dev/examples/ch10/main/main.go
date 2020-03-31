package main

import (
	"fmt"
	math "go_dev/examples/ch10/net.duokr/math"
)

func main() {
	var a, b int = 100, 200

	fmt.Println("Add demo:", math.Add(a, b))
	fmt.Println("Substract demo:", math.Subtract(a, b))
	fmt.Println("Multiply demo:", math.Multiply(a, b))
	fmt.Println("Divide demo:", math.Divide(a, b))
}
