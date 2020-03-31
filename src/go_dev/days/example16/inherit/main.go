package main

import "fmt"

type Car struct {
	Name   string
	Weight int
}

type Bike struct {
	Car
	tire int
}

type Train struct {
	c Car
}

func (car Car) run() {
	fmt.Println("Running....")
}

func (train Train) String() string {
	str := fmt.Sprintf("name=[%s], wight=[%d]", train.c.Name, train.c.Weight)
	return str
}

func main() {
	var b Bike
	b.run()

	var t Train
	t.c.Weight = 100
	t.c.Name = "Train"
	fmt.Println(t)
}
