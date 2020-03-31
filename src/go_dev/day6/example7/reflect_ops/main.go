package reflect_ops

import (
	"fmt"
	"reflect"
)

type Secret interface{}

type NotknownType struct {
	s1 string
	s2 string
	s3 string
}

func (n NotknownType) String() string {
	return n.s1 + ":" + n.s2 + ":" + n.s3
}

// var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}
var secret Secret

// var nt = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	var nt = NotknownType{"Ada", "Go", "Oberon"}
	secret = nt
	// secret = NotknownType{"Ada", "Go", "Oberon"}

	typ := reflect.TypeOf(secret)
	fmt.Println("typ: ", typ)
	value := reflect.ValueOf(secret)

	knd := value.Kind()
	fmt.Println("knd: ", knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	results := value.Method(0).Call(nil)
	fmt.Println("results: ", results)
}
