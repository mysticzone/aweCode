package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func testMap() {
	var errNotFound error = errors.New("Not Found")
	var m map[string]interface{}
	m = make(map[string]interface{})

	m["Name"] = "YanHao"
	m["Age"] = 22
	m["Score"] = 99.1

	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Error.")
	}
	fmt.Println(jsonStr)
	fmt.Printf("%s\n", string(jsonStr))
}

func main() {
	stu := Student{
		Name: "Stu01",
		Age:  20,
	}

	jsonStr, err := json.Marshal(stu)
	if err != nil {
		fmt.Printf("Error.")
	}
	fmt.Println(jsonStr)
	fmt.Printf("%s\n", string(jsonStr))

	fmt.Println("----------------------------------")
	testMap()
}
