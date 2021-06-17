package main

import (
	"fmt"
	"reflect"
)

type Secret interface {
}

type NotknownType struct {
	s1 string
	s2 string
	s3 string
}

func (n NotknownType) String() string {
	return n.s1 + ":" + n.s2 + ":" + n.s3
}

func main() {

	var nkt = NotknownType{"Ada", "Go", "Oberon"}
	var secret Secret
	secret = nkt

	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println("typ: ", typ)

	knd := value.Kind()
	fmt.Println("knd: ", knd)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	results := value.Method(0).Call(nil)
	fmt.Println("results: ", results)
}
