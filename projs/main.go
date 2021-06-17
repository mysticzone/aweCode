package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func foo() {
	var v1 byte = 'a'
	fmt.Println(v1)
	fmt.Printf("c1=%c\n", v1)

	var n1 = 10 + 'a'
	fmt.Println("n1=", n1)

	//使用的反引号 ``
	str3 := ` 
	package main
	import (
		"fmt"
		"unsafe"
	)
	
	//演示golang中bool类型使用
	func main() {
		var b = false
		fmt.Println("b=", b)
		//注意事项
		//1. bool类型占用存储空间是1个字节
		fmt.Println("b 的占用空间 =", unsafe.Sizeof(b) )
		//2. bool类型只能取true或者false
		
	}
	`
	fmt.Println(str3)

	str2 := "abc\nabc"
	fmt.Println(str2)
}

func stringDemo() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf(" b type %T n1=%v\n", b, b)

	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n =", n)
		count++
		if n == 99 {
			break
		}
	}

	fmt.Println("Count: ", count)

}

func timeDemo() {
	str2 := "hello 北京"
	//r := []rune(str2)
	//fmt.Println(r)
	//for i := 0; i < len(r); i++ {
	//	fmt.Printf("Char=%c\n", r[i])
	//}
	bytes := []byte(str2)
	fmt.Printf("%v\n", bytes)

	str := string([]byte{104, 101, 108, 108, 111, 32, 229, 140, 151, 228, 186, 172})
	fmt.Println(str)
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2

	fmt.Println("Start: ", res)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	}
	fmt.Print("n=", n)

}

func paramDemo(values []interface{}) {
	for k1, k2 := range values {
		fmt.Println(k1, k2)
	}
}

type Person struct {
	Name string
	Age  int
}

func main() {

	//var arr = []interface{}{1, 2, 3, 4, 5}
	//paramDemo(arr)
	//
	//var arr1 = []interface{}{}
	//var arr2 []int
	//fmt.Println(reflect.TypeOf(arr1))
	//fmt.Println(reflect.TypeOf(arr2))
	//arr1 = append(arr1, 100)
	//arr1 = append(arr1, 99.99)
	//arr1 = append(arr1, true)
	//arr2 = append(arr2, 100)
	////arr2 = append(arr2, 99.99)
	//fmt.Println(arr1)
	//fmt.Println(arr2)
	//test2(4)
	//timeDemo()
	//stringDemo()
	//
	//var arr01 [3]float32
	//var arr02 [3]string
	//var arr03 [3]bool
	//fmt.Printf("arr01=%v arr02=%v arr03=%v\n", arr01, arr02, arr03)

	//heroes := [...]string{"宋江", "吴用", "卢俊义"}
	////使用常规的方式遍历，我不写了..
	//
	//for i, v := range heroes {
	//	fmt.Printf("i=%v v=%v\n", i, v)
	//	fmt.Printf("heroes[%d]=%v\n", i, heroes[i])
	//}
	//
	//for _, v := range heroes {
	//	fmt.Printf("元素的值=%v\n", v)
	//}

	//fbSlice := make([]uint64, 10)
	//fbSlice[0] = 1
	//fbSlice[1] = 2

	//var intAttr [5]int = [...]int{100, 200, 300, 400, 500}
	//fmt.Println(intAttr)

	//var slice = []int{1, 2, 3, 4, 5}
	//fmt.Println(slice)
	//names := [5]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	//fmt.Println(names)

	//str := "hello"
	//arr1 := []rune(str)
	//arr1[0] = '北'
	//str = string(arr1)
	//fmt.Println(str)

	//演示二维数组的遍历
	//var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	//
	////for循环来遍历
	//for i := 0; i < len(arr3); i++ {
	//	for j := 0; j < len(arr3[i]); j++ {
	//		fmt.Printf("%v\t", arr3[i][j])
	//	}
	//	fmt.Println()
	//}

	//var p1 Person
	//p1.Age = 100
	//p1.Name = "Ming"
	//var p2 *Person = &p1
	//fmt.Println(p2.Age)
	//fmt.Println((*p2).Age)
	//p2.Name = "Tim"
	//fmt.Printf("p1的地址:%p\n", &p1)
	//fmt.Printf("p2的地址:%p p2的值%p\n", &p2, p2)

	//type A struct {
	//	Num int
	//}
	//type B struct {
	//	Num int
	//}
	//
	//type Monster struct {
	//	Name  string `json:"name"`
	//	Age   int    `json:"age"`
	//	Skill string `json:"skill"`
	//}
	//
	//var a A
	//var b B
	//a = A(b)
	//fmt.Println(a, b)
	//
	//monster := Monster{"牛魔王", 1000, "扇子"}
	//jsonStr, err := json.Marshal(monster)
	//if err != nil {
	//	fmt.Println("Json err: ", err)
	//}
	//fmt.Println("jsonStr: ", string(jsonStr))

	//type Stu struct {
	//	Name string
	//	Age  int
	//}
	//
	//var stu5 *Stu = &Stu{"xiaowang", 99}
	//stu6 := &Stu{"xiaowang", 100}
	//var stu7 = &Stu{
	//	Name: "xiaoqiang",
	//	Age:  98,
	//}
	//fmt.Println(*stu5, *stu6, *stu7)

	//map1 := make(map[int]int, 2)
	//map1[1] = 90
	//map1[2] = 88
	//map1[10] = 1
	//map1[20] = 2
	//fmt.Println(map1)

	//users := make(map[string]map[string]string, 10)
	//users["smith"] = make(map[string]string, 2)
	//users["smith"]["pwd"] = "999999"
	//users["smith"]["nickname"] = "小花猫"
	//fmt.Println(users)

	//map1 := make(map[int]int, 10)
	//map1[10] = 100
	//map1[1] = 13
	//map1[4] = 56
	//map1[8] = 90
	//map1[2] = 200
	//
	//fmt.Println(map1)
	//var keys []int
	//for k, _ := range map1 {
	//	keys = append(keys, k)
	//}
	////排序
	//sort.Ints(keys)
	//fmt.Println(keys)
	//fmt.Println(map1)
	//
	//for _, k := range keys {
	//	fmt.Printf("map1[%v]=%v \n", k, map1[k])
	//}

	//var monsters []map[string]string
	//monsters = make([]map[string]string, 2) //准备放入两个妖怪
	////2. 增加第一个妖怪的信息
	//if monsters[0] == nil {
	//	monsters[0] = make(map[string]string, 2)
	//	monsters[0]["name"] = "牛魔王"
	//	monsters[0]["age"] = "500"
	//}
	//
	//if monsters[1] == nil {
	//	monsters[1] = make(map[string]string, 2)
	//	monsters[1]["name"] = "玉兔精"
	//	monsters[1]["age"] = "400"
	//}
	//fmt.Println(monsters)
}
