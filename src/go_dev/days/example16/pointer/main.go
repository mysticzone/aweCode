package main

import "fmt"

type Student struct {
	Name  string
	Age   int32
	Score float32
	next  *Student
}

func trans(p *Student) {
	// var p *Student = &head
	for p != nil {
		fmt.Println(p)
		p = p.next
	}
}

func inertTail(p *Student) {
	for i := 0; i <= 10; i++ {
		var stu Student
		stu.Name = "hello"
		stu.Age = int32(i)
		stu.Score = float32(i)

		p.next = &stu
		p = &stu
	}
}

func inertHead(p **Student) {
	var stu Student
	for i := 0; i <= 10; i++ {
		stu.Name = "hello"
		stu.Age = int32(i)
		stu.Score = float32(i)

		stu.next = *p
		*p = &stu
	}
	fmt.Println("> ", p)
}

func main() {
	var head *Student = &Student{}
	head.Name = "head123"
	head.Age = 10
	head.Score = 10

	// var stu1 Student
	// stu1.Name = "stu1111"
	// stu1.Age = 11
	// stu1.Score = 11
	// // stu1.next = nil
	// head.next = &stu1
	//
	// var stu2 Student
	// stu2.Name = "stu2222"
	// stu2.Age = 22
	// stu2.Score = 22
	// // stu1.next = nil
	// stu1.next = &stu2

	// var tail = &head
	// inertHead(head)

	fmt.Println(head)
	fmt.Println(&head)
	fmt.Println("-------------")
	inertHead(&head)
	// inertTail(&head)
	trans(head)

	// fmt.Println(head)
	// fmt.Printf("%p", &stu1)
	// fmt.Println(stu1.next)

	// for ; p.next != nil; p = p.next {
	// 	fmt.Println(p.Name)
	// }
}
