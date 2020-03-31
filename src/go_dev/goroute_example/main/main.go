package main

import (
	"fmt"
	"go_dev/goroute_example/goroute"
)


func main(){
	var pipe chan int
	pipe = make(chan int, 2)
	go goroute.Add(2, 2, pipe)
	
	sum :=<- pipe
	fmt.Println("Sum: ", sum)
}
