package main

import "fmt"

type Test interface {
	Print()
	Sleep()
}

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (stu Student) Print() {
	fmt.Println(stu)
}

func (stu Student) Sleep() {
	fmt.Println("Sleeping....")
}

func main() {
	var t Test
	t.Sleep()
	fmt.Println("t: ", t)
	var stu Student
	stu.Name = "yh"
	stu.Age = 100
	stu.Score = 101

	// stu.Print()
	fmt.Println("stu: ", stu)
	t = stu
	fmt.Println("t: ", t)
	t.Print()
	t.Sleep()
}
