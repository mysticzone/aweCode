package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func get_sum_of_divisible(num int, divider int, resultChan chan int) {
	sum := 0

	for value := 0; value < num; value++ {
		if value%divider == 0 {
			sum += value
		}
	}
	resultChan <- sum
}

func list_elem(n int, tag string) {
	for i := 0; i < n; i++ {
		fmt.Println(tag, " >> ", i)

		tick := time.Duration(rand.Intn(100))
		time.Sleep(time.Millisecond * tick)
	}
}

func fixed_shotting(msg_chan chan string) {
	var times = 3
	var t = 1
	for {
		if t <= times {
			msg_chan <- "Fixed shotting"
		}
		t++
		time.Sleep(time.Second * 1)
	}
}

func three_point_shotting(msg_chan chan string) {
	var times = 5
	var t = 1
	for {
		if t <= times {
			msg_chan <- "three point shotting"
		}
		t++
		time.Sleep(time.Second * 1)
	}
}

func count(msg_chan chan string) {
	for {
		msg := <-msg_chan
		fmt.Println("---->", msg)
		time.Sleep(time.Second * 1)
	}
}

func shotting(msg_chan chan string) {
	var group = 1
	for {
		for i := 0; i <= 10; i++ {
			msg_chan <- strconv.Itoa(group) + ":" + strconv.Itoa(i)
		}
		group++
		time.Sleep(time.Second * 10)
	}
}

func count_shot(msg_chan chan string) {
	for {
		fmt.Println(<-msg_chan)
	}
}

func main() {
	LIMIT := 10
	resultChan := make(chan int, 10)

	t_start := time.Now()
	go get_sum_of_divisible(LIMIT, 3, resultChan)
	go get_sum_of_divisible(LIMIT, 5, resultChan)

	sum3, sum5 := <-resultChan, <-resultChan

	go get_sum_of_divisible(LIMIT, 15, resultChan)
	sum15 := <-resultChan

	sum := sum3 + sum5 - sum15
	t_end := time.Now()
	fmt.Println("Sum: ", sum)
	fmt.Println("Cost: ", t_end.Sub(t_start))

	fmt.Println("------------------------------")
	// go list_elem(10, "A")
	// go list_elem(20, "B")

	// var name string
	// fmt.Scanln(&name)
	fmt.Println("------------------------------")

	// var c chan string
	// // var c3_point chan string
	// c = make(chan string)
	// c3_point := make(chan string)
	// go fixed_shotting(c)
	// go three_point_shotting(c3_point)
	// // go count(c)
	//
	// go func() {
	// 	for {
	// 		select {
	// 		case msg1 := <-c:
	// 			fmt.Println(msg1)
	// 		case msg2 := <-c3_point:
	// 			fmt.Println(msg2)
	// 		case <-time.After(time.Second * 20):
	// 			fmt.Println("Timeout, check again...")
	// 		}
	// 	}
	// }()

	var cc chan string
	cc = make(chan string, 20)
	go shotting(cc)
	go count_shot(cc)
	var name string
	fmt.Scanln(&name)

}
