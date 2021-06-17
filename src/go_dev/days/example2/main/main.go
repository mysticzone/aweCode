package main

import "fmt"

func output(n int){
	for i :=0; i <= n; i++ {
		fmt.Printf("%d + %d = %d\n", i, n - i, n)
	}
}

func main(){
	output(10)
}