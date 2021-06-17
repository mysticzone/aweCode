package demo

import "fmt"


/*
水仙花数
*/
func Narcissistic(){
	var n int
	_, _ = fmt.Scanf("%d", &n)

}




/*
101-200之间的素数
*/
func Prime(){
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println("n: ", n)

	var f bool = true
	for i := 2; i < n/2; i++{
		if n % i == 0 {
			f = false
			break
		}
	}
	if f {
		fmt.Println("素数：", n)
	}
}