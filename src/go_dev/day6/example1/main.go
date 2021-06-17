package main

import "fmt"

type Test interface {
	GetName() string
	Run()
	DiDi()
}

type Cars struct {
	Name string
}

func (c Cars) GetName() string {
	return c.Name
}

func (c Cars) Run() {
	fmt.Printf("%s is running.\n", c.Name)
}

func (c Cars) DiDi() {
	fmt.Printf("%s is didi.\n", c.Name)
}

type BYD struct {
	Name string
}

func (c BYD) GetName() string {
	return c.Name
}

func (c BYD) Run() {
	fmt.Printf("%s is running.\n", c.Name)
}

func (c BYD) DiDi() {
	fmt.Printf("%s is didi.\n", c.Name)
}

func main() {
	var t Test
	var c Cars
	c.Name = "BWM"

	t = c
	// fmt.Println(t.GetName())
	t.Run()

	var b BYD
	b.Name = "BYD"
	t = b
	t.Run()

	// fmt.Printf("%T", t)
}
