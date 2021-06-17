package main

import (
	"fmt"
	// "testing"
)

func multiply(x string, y int) {
	result := ""
	for i := 0; i < y; i++ {
		result = fmt.Sprintf("%s%s", result, x)
		fmt.Println("> : ", result)
	}
	fmt.Println("result: ", result)
}
func main() {
	multiply("a", 5)
}
