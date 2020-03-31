package add

import "fmt"
import "time"

var Name string
var Age int

const (
	a = iota
	b
	c
)

func Test(){
	Name = "Hello World"
	Age = 10

}

func Calc(){

	const (
		Man = 1
		Female = 2
	)
	//var Man int = 1
	//var Female int = 2

    for {
		seconds := int(time.Now().Unix())
		fmt.Println("Seconds: ", seconds)

		ret := seconds % Female
		if ret == 0 {
			fmt.Println("Female: ", Female)
		} else {
			fmt.Println("Man: ", Man)
		}
		time.Sleep(time.Millisecond*1000)
	}

}

func init(){
	fmt.Println("init....")
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)
	fmt.Println("C: ", c)
}

