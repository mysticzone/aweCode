package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (stu Student) Print() {
	fmt.Println(stu)
}

func Test(b interface{}) {
	// fmt.Println(reflect.TypeOf(b))
	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}
}

func TestStruct(b interface{}) {
	val := reflect.ValueOf(b)
	k := val.Kind()
	if k != reflect.Struct {
		fmt.Println("expect struct")
	}
	fmt.Println("Fields:", val.NumField())
	fmt.Println("Methods: ", val.NumMethod())

}

type Stringer interface {
	String() string
}

func main() {
	// var a int
	var stu = Student{Name: "YangHao"}
	// Test(a)
	Test(stu)
	// fmt.Println(stu)
	TestStruct(stu)
	fmt.Println("----------------------------------")
	var t int
	var x interface{}
	x = t
	y, ok := x.(int)
	fmt.Println("Y: ", y, "OK: ", ok)
	fmt.Println("----------------------------------")

	var a int
	var b interface{}
	b = a
	fmt.Println("A: ", a, "B: ", b)
	fmt.Println("----------------------------------")

}
