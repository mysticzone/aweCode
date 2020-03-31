package main

import "fmt"

type integer int

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (p *Student) init(name string, age int, score float32) {
	p.Name = name
	p.Age = age
	p.Score = score
}

func (p Student) get() Student {
	return p
}

func (p integer) get() {
	fmt.Println(p)
}

func (p *integer) set(b integer) {
	*p = b
	fmt.Println(p)
}

func main() {
	var stu Student
	stu.init("yang", 110, 112.1)
	fmt.Println(stu)
	fmt.Println(stu.get())

	var a integer
	a = 100
	a.get()
	a.set(101)
	a.get()

}
