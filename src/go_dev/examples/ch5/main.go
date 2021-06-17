package main

import "fmt"

func Test1() {
	var x = make([]float64, 5)
	fmt.Printf("capcity: %d length: %d\n", cap(x), len(x))

	var y = make([]float64, 5, 10)
	fmt.Printf("capcity: %d length: %d\n", cap(y), len(y))

	for i := 0; i < len(x); i++ {
		x[i] = float64(i)
	}
	fmt.Println("X: ", x)

	for i := 0; i < len(y); i++ {
		y[i] = float64(i)
	}
	fmt.Println("Y: ", y)
}

func Test2() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("{1:3}", arr[1:3])
}

func Test3() {
	var arr1 = make([]int, 5, 10)
	fmt.Println("arr1: ", arr1)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Println("arr1: ", arr1)

	arr1 = append(arr1, 5, 6, 7, 8)
	fmt.Println("arr1: ", arr1)
	fmt.Printf("Capacity: %d  Length: %d\n", cap(arr1), len(arr1))

}

func Test4() {
	var arr1 = []int{1, 2, 3, 4, 5, 6}
	var slice1 = make([]int, 5, 10)

	copy(slice1, arr1)
	fmt.Println(slice1)
}

func Test5() {
	var x = map[string]string{
		"A": "Apple",
		"P": "Pear",
		"O": "Orange",
	}
	// fmt.Println("X: ", x)
	for k, v := range x {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
}

func Test6() {
	var x map[string]string
	x = make(map[string]string)
	fmt.Printf("Length: %d\n", len(x))

	x["A"] = "Apple"
	x["O"] = "Orange"
	x["P"] = "pear"

	fmt.Printf("Length: %d\n", len(x))

	for k, v := range x {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}

	fmt.Printf("C: %d", x["C"])

}

func Test7() {
	var x = make(map[string]int)
	x["A"] = 20
	x["O"] = 0
	x["P"] = 100

	for k, v := range x {
		fmt.Printf("Key: %s Value: %d\n", k, v)
	}

	v, o := x["A"]
	fmt.Println("v: ", v, "o: ", o)

	if val, ok := x["A"]; ok {
		fmt.Printf("OK: %s Value: %d\n", ok, val)
	}
	fmt.Println(x)

	delete(x, "C")
	fmt.Println(x)
}

func Test8() {
	var fb = make(map[string]map[string]int)
	fb["123"] = map[string]int{"Jemy": 25}
	fb["124"] = map[string]int{"Tom": 26}
	fb["125"] = map[string]int{"Lilei": 18}
	// fb["126"] = {s := "Andy": 19}

	for no, info := range fb {
		fmt.Println("No: ", no)
		for name, age := range info {
			fmt.Printf("{Name: %s, Age: %d}\n", name, age)
		}
	}
}

func main() {
	// Test3()
	// Test4()
	// Test5()
	// Test7()
	Test8()
}
