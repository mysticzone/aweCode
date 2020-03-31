package main


import (
	"fmt"
	"math/rand"
)


var c byte = 'a'

func rand_num(n int){
	for i := 0; i< n; i++ {
		temp_int := rand.Intn(100)
		fmt.Println(temp_int)
	}

}

func rand_float(n int){
	for i := 0; i< n; i++ {
		temp_int := rand.Float32()
		fmt.Println(temp_int)
	}
}

func reserve(str string) string {
	var result []byte
	temp := []byte(str)
	length := len(str)

	for i :=0; i < length; i++ {
		result = append(result, temp[length-i-1])
	}
	return string(result)
}

func main(){
	rand_num(10)
	rand_float(10)
	fmt.Println(c)

	var s string = `hello \n world`
	fmt.Println(s)


}
