package main

import (
	"fmt"
	"os"
)

func main(){
	goos := os.Getenv("os")

	fmt.Println("goos: ", goos)
	fmt.Println("Env:", os.Environ())
}