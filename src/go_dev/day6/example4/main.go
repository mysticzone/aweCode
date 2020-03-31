package main

import "fmt"

type Student struct {
	Name string
	Sex  string
}

func Test(x interface{}) {
	y, ok := x.(int)
	if !ok {
		fmt.Println("converted fail")
	} else {
		fmt.Println(y + 4)
	}
}
func classifier(items ...interface{}) {
	for index, value := range items {
		switch value.(type) {
		case bool:
			fmt.Printf("%d params is bool, value is %v\n", index, value)
		case int, int32, int64:
			fmt.Printf("%d params is int32, value is %v\n", index, value)
		case float32, float64:
			fmt.Printf("%d params is float32, value is %v\n", index, value)
		case string:
			fmt.Printf("%d params is string, value is %v\n", index, value)
		case Student:
			fmt.Printf("%d params is Student, value is %v\n", index, value)
		}
	}
}
func main() {
	// var x string
	var stu Student
	Test(stu)
	classifier(10, 2.2, "hello world", true, stu)

}
