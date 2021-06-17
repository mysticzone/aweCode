package main

import (
	"fmt"
)

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
}

func (p *File) Read() {
	fmt.Println("read data")
}

func (p *File) Write() {
	fmt.Println("write data")
}

func TestRW(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	p := File{}

	// p.Read()

	TestRW(&p)

	var a interface{}
	var b int

	a = b
	c := a.(int)
	fmt.Printf("%d %T\n", c, c)

	var f interface{}
	f = p
	if v, ok := f.(ReadWriter); ok {
		fmt.Println(v, ok)
	}
}
