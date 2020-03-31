package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"stu_name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {

	var stu Student = Student{
		Name:  "yang",
		Age:   110,
		Score: 111,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("Json encode failed. err: ", err)
		return
	}

	fmt.Println(string(data))

}
