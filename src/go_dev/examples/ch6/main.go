package main

import (
	"bufio"
	"fmt"
	"os"
)

func Test_sum(base int, arr ...int) {
	sum := base
	for _, v := range arr {
		sum += v
	}
	fmt.Println("Sum: ", sum)
}

var b = func(base int, arr ...int) {
	sum := base
	for _, v := range arr {
		sum += v
	}
	fmt.Println("Sum: ", sum)
}

// func createEvenGenerator() func() unit {
// 	i := unit(0)
//
// 	return func() (retVal unit) {
// 		retVal = i
// 		i += 2
// 		return
// 	}
// }

func first() {
	fmt.Println("first func run")
}

func second() {
	fmt.Println("second func run")
}

func Reader() {
	fname := "/Users/yanghao/default.log"
	f, err := os.Open(fname)
	defer f.Close()

	if err != nil {
		os.Exit(1)
	}

	bReader := bufio.NewReader(f)
	for {
		line, ok := bReader.ReadString('\n')
		if ok != nil {
			break
		}
		fmt.Printf("%s", line)
	}
}

func TestRecover() {
	defer func() {
		msg := recover()
		fmt.Printf("> %s", msg)
	}()
	fmt.Println("I am walking and singing...")
	panic("It starts to rain cats and dogs")
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	Test_sum(100, arr...)
	b(200, arr...)

	// nextEven := createEvenGenerator()
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())

	defer second()
	first()

	// Reader()

	TestRecover()
}
