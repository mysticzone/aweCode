package example1
import (
	"fmt"
	//"time"
)

var pipe chan int

func add(a int, b int, c chan int){
	var sum int
	sum = a + b
	//return sum
	c <- sum
}

func main(){
	var c int
	//c = add(100, 1)
	//go test_goroute(200, 2)
	fmt.Println("Sum: ", c)
	fmt.Println("Hello World")
	//for i := 0; i < 100; i++ {
	//	go test_print(i)
	//	time.Sleep(10)
	//}

	//test_pipe()
	pipe = make(chan int, 3)
	go add(2, 5, pipe)
	sum :=<- pipe
	fmt.Println("sum=", sum)

}