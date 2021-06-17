package main

import (
	"fmt"
)

type Phone interface {
	Call()
	Sales() int
}

type NokiaPhone struct {
	Name  string
	Price int
}

type IPhone struct {
	Name  string
	Price int
}

type Person struct {
	phones []Phone
	name   string
	age    int
}

func (phone NokiaPhone) Call() {
	fmt.Println("Nokia Calling...")
}

func (phone IPhone) Call() {
	fmt.Println("IPhone Calling...")
}

func (phone NokiaPhone) Sales() int {
	return phone.Price
}

func (phone IPhone) Sales() int {
	return phone.Price
}

func (person Person) TotalCost() int {
	var sum = 0
	for _, ph := range person.phones {
		sum += ph.Sales()
	}
	return sum
}

func main() {
	var phone Phone
	fmt.Println("phone: ", &phone)

	var nokia = NokiaPhone{Name: "Nokina"}
	phone = nokia
	// phone = new(NokiaPhone)
	fmt.Println("phone: ", phone)

	fmt.Println("phone: ", &phone)
	phone.Call()

	var iphone = IPhone{Name: "IPhone"}
	phone = iphone
	// phone = new(IPhone)
	fmt.Println("phone: ", phone)
	fmt.Println("phone: ", &phone)
	phone.Call()

	fmt.Println("-------------------------------------")

	var phs = [5]Phone{
		NokiaPhone{Price: 350},
		IPhone{Price: 5000},
		IPhone{Price: 3400},
		NokiaPhone{Price: 450},
		IPhone{Price: 5000},
	}

	TotalCost := 0
	for _, ph := range phs {
		fmt.Println("ph: ", ph)
		TotalCost += ph.Sales()
	}
	fmt.Println("TotalCost: ", TotalCost)
	fmt.Println("-------------------------------------")

	var person = Person{
		name:   "Jemy",
		age:    25,
		phones: phs[:],
	}
	fmt.Println("Person Name: ", person.name)
	fmt.Println("Person Age: ", person.age)
	fmt.Println("Total Cost: ", person.TotalCost())
}
