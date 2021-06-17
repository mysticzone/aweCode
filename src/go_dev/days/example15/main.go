package main

import "fmt"

type Student struct {
	name  string
	age   int64
	score float32
}

func main() {
	var stu Student
	stu.age = 18
	stu.name = "yh"
	stu.score = 99.1

	var stu1 Student = Student{
		name: "hao",
		age:  0,
	}

	var stu2 = Student{
		name:  "yang",
		score: 0,
	}

	fmt.Println(stu1)
	fmt.Println(stu2)

	fmt.Println(stu)
	fmt.Printf("%p, %p, %p, %p", &stu, &stu.name, &stu.age, &stu.score)
}
