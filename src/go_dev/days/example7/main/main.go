package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func modify(p *int) {
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)
	*p = 123
}

type func_type func(int, int) int

func add(x int, y int) int {
	return x + y
}

func operator(op func_type, a int, b int) int {
	return op(a, b)
}

func main() {
	var s string = "http://www.baidu.com, www.baidu.com"
	fmt.Println(strings.HasPrefix(s, "http"))
	fmt.Println(strings.Replace(s, "baidu", "google", 3))
	fmt.Println("Time: ", time.Now())
	fmt.Println(time.Time{})
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04"))

	fmt.Println("-------------------------------------")

	var x int = 'a'
	fmt.Println("x: ", x)
	fmt.Println("&x: ", &x)

	var p *int
	p = &x
	//_, _ = fmt.Scanf()
	fmt.Println("&p: ", &p)
	fmt.Println("p", p)
	fmt.Println("*p: ", *p)

	*p = 'b'
	fmt.Println("x: ", x)
	fmt.Println("&x: ", &x)

	fmt.Println("-------------------------------------")
	modify(&x)
	fmt.Println("x: ", x)

	var b int = 0
	switch b {
	case 0:
		fmt.Println("sss")
	case 1:
		fmt.Println("wwww")
	default:
		fmt.Println("ddd")
	}

	fmt.Println("rand: ", rand.Intn(100))

	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%s", "A")
		}
		fmt.Println()
	}

	var ss string = "aaa bbb ccc  world, 中国"
	for i, v := range ss {
		fmt.Println(i, v, len([]byte(string(v))))
	}

	fmt.Println("-------------------------------------")

	var u int = 10
	var v int = 20
	c := add
	fmt.Println(c)
	ret := operator(c, u, v)
	fmt.Println("ret: ", ret)
}
