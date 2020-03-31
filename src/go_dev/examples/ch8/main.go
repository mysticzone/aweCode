package main

import "fmt"

type Rect struct {
	width  float64
	height float64
}

func StructRect() {
	var rect Rect
	rect.height = 100
	rect.width = 200
	fmt.Println(rect.width * rect.height)

	var r = Rect{width: 200, height: 200}
	fmt.Println(r.width * r.height)

}

func double_area(rect Rect) float64 {
	rect.width *= 2
	rect.height *= 3
	return rect.width * rect.height
}

func StructArea() {
	var rect = Rect{width: 10, height: 20}
	area := double_area(rect)
	fmt.Println("Area: ", area)
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func StructMethodArea() {
	var rect = Rect{width: 10, height: 20}
	fmt.Println("Area: ", rect.Area())
}

func (r *Rect) PtrArea() float64 {
	r.height = 22
	return r.width * r.height
}

func StructPtrArea() {
	var rect = new(Rect)
	// var rect = Rect{width: 10, height: 20}
	rect.width = 30
	rect.height = 20
	fmt.Println(rect)
	fmt.Println("Area: ", rect.PtrArea())
	fmt.Println(rect)
}

func (r *Rect) DoubleArea() float64 {
	r.height *= 2
	r.width *= 2
	return r.width * r.height
}

func StructPtrDoubleArea() {
	// var rect = new(Rect)
	var rect = Rect{width: 10, height: 20}
	rect.width = 30
	rect.height = 20
	fmt.Println(rect)
	fmt.Println("Area: ", rect.DoubleArea())
	fmt.Println(rect)
}

type Phonei struct {
	price int
	color string
}

func (phone Phonei) Ringing() {
	fmt.Println("Ringing....")
}

type Iphone struct {
	Phonei
	model string
}

func IphoneFunc() {
	var p Iphone
	p.Ringing()
}

type NokiaPhone struct {
	price int
}

type IPhone struct {
	price int
}

type Phone interface {
	Call()
	Sales() int
}

type Person struct {
	phones []Phone
	name   string
	age    int
}

func (phone NokiaPhone) Call() {
	fmt.Println("I am Nokia, I can call you!....")
}

func (phone NokiaPhone) Sales() int {
	return phone.price
}

func (phone IPhone) Call() {
	fmt.Println("I am iPhone, I can call you!")
}

func (phone IPhone) Sales() int {
	return phone.price
}

func (person Person) totalCost() int {
	var sum = 0
	for _, person := range person.phones {
		sum += person.Sales()
	}
	return sum
}

func CallPhone() {
	// var nokia NokiaPhone
	// var iphone IPhone
	// var phone Phone
	//
	// phone = new(NokiaPhone)
	// phone.Call()
	// phone = new(IPhone)
	// phone.Call()
	// nokia.Call()
	// iphone.Call()

	var phs = [5]Phone{
		NokiaPhone{price: 350},
		IPhone{price: 5000},
		IPhone{price: 3400},
		NokiaPhone{price: 450},
		IPhone{price: 5000},
	}
	var totalSales = 0

	for _, ph := range phs {
		totalSales += ph.Sales()
	}
	fmt.Println("TotalSales: ", totalSales)

	fmt.Println("---------------------------")

	var person = Person{name: "Jemy", age: 25, phones: phs[:]}
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.totalCost())

}

func main() {
	// StructRect()
	//
	// StructArea()
	//
	// StructMethodArea()

	// StructPtrArea()
	// StructPtrDoubleArea()

	// IphoneFunc()
	CallPhone()
}
