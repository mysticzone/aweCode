package example1

import "fmt"

func test_pipe(){
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	pipe <- 4


	fmt.Println("Length: ", len(pipe))
}
