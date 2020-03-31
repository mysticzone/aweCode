package main

import (
	"fmt"
	"go_dev/days/example3/add"
)

func main() {
	add.Test()
	add.Calc()
	fmt.Println("Name: ", add.Name)
	fmt.Println("age: ", add.Age)
}
