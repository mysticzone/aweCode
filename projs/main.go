package main

func main() {
}

//func main() {
//
//	conn, err := redis.Dial("tcp", "192.168.146.54:6379")
//	if err != nil {
//		panic(err)
//	}
//	defer func(conn redis.Conn) {
//		_ = conn.Close()
//	}(conn)
//	fmt.Println(conn)
//	r, err := conn.Do("set", "k1", "hello")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(r.(string))
//}

//func main() {
//	var str string = "hello"
//	fs := reflect.ValueOf(&str)
//	fmt.Println(fs.Kind(), fs.Type())
//	x := "world"
//	fs.Elem().SetString(x)
//	fmt.Printf("%v\n", str)
//	//fmt.Println(fs.Elem())
//}

//func reflect01(b interface{}) {
//	rVal := reflect.ValueOf(b)
//	fmt.Println(rVal.Kind())
//	//rVal.SetInt(20)
//	fmt.Println(rVal.Elem())
//	rVal.Elem().SetInt(20)
//}
//
//func main() {
//	var num int = 10
//	reflect01(&num)
//	fmt.Println(num)
//}

//import (
//	"fmt"
//	"reflect"
//)
//
//type Student struct {
//	Name string
//	Age  int
//}
//
//func reflectTest2(b interface{}) {
//	rType := reflect.TypeOf(b)
//	fmt.Println("rType: ", rType)
//
//	rVal := reflect.ValueOf(b)
//	fmt.Printf("rVal:%v\trVal Type: %T\n", rVal, rVal)
//
//	rInf := rVal.Interface()
//	fmt.Printf("rInf:%v\trInf Type: %T\n", rInf, rInf)
//}
//
//func main() {
//	stu := Student{
//		Name: "Tom",
//		Age:  20,
//	}
//	reflectTest2(stu)
//}

//func reflectTest1(b interface{}) {
//	rType := reflect.TypeOf(b)
//	fmt.Println("rType: ", rType)
//
//	rVal := reflect.ValueOf(b)
//	fmt.Printf("rVal: %v\trVal Type: %T\n", rVal, rVal)
//
//	iv := rVal.Interface()
//	fmt.Printf("iv: %v\tiv Type: %T\n", iv, iv)
//
//	num := iv.(int)
//	fmt.Printf("num: %v\tnum Type: %T\n", num, num)
//}
//
//func main() {
//	var num int = 100
//	reflectTest1(num)
//}

//import (
//	"fmt"
//)
//
//type Animal interface {
//	Speak() string
//}
//
//type Dog struct {
//}
//
//func (d Dog) Speak() string {
//	return "Woof!"
//}
//
//type Cat struct {
//}
//
////1
//func (c *Cat) Speak() string {
//	return "Meow!"
//}
//
//type Llama struct {
//}
//
//func (l Llama) Speak() string {
//	return "?????"
//}
//
//type JavaProgrammer struct {
//}
//
//func (j JavaProgrammer) Speak() string {
//	return "Design patterns!"
//}
//
//func main() {
//	animals := []Animal{Dog{}, &Cat{}, Llama{}, JavaProgrammer{}}
//	for _, animal := range animals {
//		fmt.Println(animal.Speak())
//	}
//}

//import "fmt"
//
////1
//type I interface {
//	Get() int
//	Set(int)
//}
//
////2
//type S struct {
//	Age int
//}
//
//func (s S) Get() int {
//	return s.Age
//}
//
//func (s *S) Set(age int) {
//	s.Age = age
//}
//
////3
//func f(i I) {
//	i.Set(10)
//	fmt.Println(i.Get())
//}
//
//func main() {
//	s := S{}
//	f(&s) //4
//}

////定义了一个Monster结构体
//type Monster struct {
//	Name  string  `json:"name"`
//	Age   int     `json:"monster_age"`
//	Score float32 `json:"成绩"`
//	Sex   string
//}
//
////方法，返回两个数的和
//func (s Monster) GetSum(n1, n2 int) int {
//	return n1 + n2
//}
//
////方法， 接收四个值，给s赋值
//func (s Monster) Set(name string, age int, score float32, sex string) {
//	s.Name = name
//	s.Age = age
//	s.Score = score
//	s.Sex = sex
//}
//
////方法，显示s的值
//func (s Monster) Print() {
//	fmt.Println("---start~----")
//	fmt.Println(s)
//	fmt.Println("---end~----")
//}
//
//func TestStruct(a interface{}) {
//	//获取reflect.Type 类型
//	typ := reflect.TypeOf(a)
//	//获取reflect.Value 类型
//	val := reflect.ValueOf(a)
//	//获取到a对应的类别
//	kd := val.Kind()
//	//如果传入的不是struct，就退出
//	if kd != reflect.Struct {
//		fmt.Println("expect struct")
//		return
//	}
//
//	//获取到该结构体有几个字段
//	num := val.NumField()
//
//	fmt.Printf("struct has %d fields\n", num) //4
//	//变量结构体的所有字段
//	for i := 0; i < num; i++ {
//		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
//		//获取到struct标签, 注意需要通过reflect.Type来获取tag标签的值
//		tagVal := typ.Field(i).Tag.Get("json")
//		//如果该字段于tag标签就显示，否则就不显示
//		if tagVal != "" {
//			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
//		}
//	}
//
//	//获取到该结构体有多少个方法
//	numOfMethod := val.NumMethod()
//	fmt.Printf("struct has %d methods\n", numOfMethod)
//
//	//var params []reflect.Value
//	//方法的排序默认是按照 函数名的排序（ASCII码）
//	val.Method(1).Call(nil) //获取到第二个方法。调用它
//
//	//调用结构体的第1个方法Method(0)
//	var params []reflect.Value //声明了 []reflect.Value
//	params = append(params, reflect.ValueOf(10))
//	params = append(params, reflect.ValueOf(40))
//	fmt.Println("params: ", params)
//	res := val.Method(0).Call(params) //传入的参数是 []reflect.Value, 返回[]reflect.Value
//	fmt.Println("res=", res[0].Int()) //返回结果, 返回的结果是 []reflect.Value*/
//
//}
//func main() {
//	//创建了一个Monster实例
//	var a Monster = Monster{
//		Name:  "黄鼠狼精",
//		Age:   400,
//		Score: 30.8,
//	}
//	//将Monster实例传递给TestStruct函数
//	TestStruct(a)
//}

//func main() {
//	//var wCh chan<- int
//	//wCh = make(chan int, 1)
//	//wCh <- 1
//
//	intCh := make(chan int, 10)
//	for i := 0; i < 10; i++ {
//		intCh <- i
//	}
//	strCh := make(chan string, 5)
//	for i := 0; i < 5; i++ {
//		strCh <- "hello" + fmt.Sprintf(" %d", i)
//	}
//
//	for {
//		select {
//		case v := <-intCh:
//			fmt.Println("int: ", v)
//		case v := <-strCh:
//			fmt.Println("str: ", v)
//		default:
//			fmt.Println("没有找到了")
//			return
//		}
//	}
//
//}

//
////专门演示反射
//func reflectTest01(b interface{}) {
//
//	//通过反射获取的传入的变量的 type , kind, 值
//	//1. 先获取到 reflect.Type
//	rTyp := reflect.TypeOf(b)
//	fmt.Println("rType=", rTyp)
//
//	//2. 获取到 reflect.Value
//	rVal := reflect.ValueOf(b)
//
//	n2 := 2 + rVal.Int()
//	//n3 := rVal.Float()
//	fmt.Println("n2=", n2)
//	//fmt.Println("n3=", n3)
//
//	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)
//
//	//下面我们将 rVal 转成 interface{}
//	iV := rVal.Interface()
//	//将 interface{} 通过断言转成需要的类型
//	num2 := iV.(int)
//	fmt.Println("num2=", num2)
//}
//
////专门演示反射[对结构体的反射]
//func reflectTest02(b interface{}) {
//
//	//通过反射获取的传入的变量的 type , kind, 值
//	//1. 先获取到 reflect.Type
//	rTyp := reflect.TypeOf(b)
//	fmt.Println("rType=", rTyp)
//
//	//2. 获取到 reflect.Value
//	rVal := reflect.ValueOf(b)
//
//	//3. 获取 变量对应的Kind
//	//(1) rVal.Kind() ==>
//	kind1 := rVal.Kind()
//	//(2) rTyp.Kind() ==>
//	kind2 := rTyp.Kind()
//	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)
//
//	//下面我们将 rVal 转成 interface{}
//	iV := rVal.Interface()
//	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
//	//将 interface{} 通过断言转成需要的类型
//	//这里，我们就简单使用了一带检测的类型断言.
//	//同学们可以使用 swtich 的断言形式来做的更加的灵活
//	stu, ok := iV.(Student)
//	if ok {
//		fmt.Printf("stu.Name=%v\n", stu.Name)
//	}
//}
//
//type Student struct {
//	Name string
//	Age  int
//}
//
//func main() {
//	//1. 先定义一个int
//	var num int = 100
//	reflectTest01(num)
//
//	//2. 定义一个Student的实例
//	stu := Student{
//		Name: "tom",
//		Age:  20,
//	}
//	reflectTest02(stu)
//}

//var flag = false
//
//func writeData(ch chan<- int) {
//	for i := 0; i <= 50; i++ {
//		fmt.Println("write: ", i)
//
//		ch <- i
//		//time.Sleep(time.Second)
//	}
//	close(ch)
//}
//
//func readeData(ch <-chan int) {
//	for ret := range ch {
//		fmt.Println("read: ", ret)
//	}
//	flag = true
//}
//
//func main() {
//	ch := make(chan int, 50)
//
//	go writeData(ch)
//	go readeData(ch)
//	for {
//		if flag {
//			break
//		}
//	}
//
//}

//func main() {
//	var nums = []int{1, 2, 3, 4, 5, 6, 7}
//	fmt.Println(nums)
//
//}

//func reflectFunc(b interface{}) {
//	rVal := reflect.ValueOf(b)
//	fmt.Println("Kind ", rVal.Kind())
//	fmt.Printf("rVal kind=%v\n", rVal.Kind())
//	rVal.Elem().SetInt(20)
//
//}
//
//func main() {
//	var num int = 10
//	reflectFunc(&num)
//	fmt.Println("num: ", num)
//
//}

//func putNum(intChan chan int) {
//	for i := 0; i <= 8000; i++ {
//		intChan <- i
//	}
//	close(intChan)
//}
//
//func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
//	var flag bool
//	for {
//		num, ok := <-intChan
//		if !ok {
//			break
//		}
//
//		flag = true
//		for i := 2; i < num; i++ {
//			if num%i == 0 {
//				flag = false
//				break
//			}
//		}
//
//		if flag {
//			primeChan <- num
//		}
//	}
//	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
//
//	exitChan <- true
//}
//
//func main() {
//	intChan := make(chan int, 1000)
//	primeChan := make(chan int, 20000)
//	exitChan := make(chan bool, 8)
//
//	start := time.Now().Unix()
//
//	go putNum(intChan)
//
//	for i := 0; i < 8; i++ {
//		go primeNum(intChan, primeChan, exitChan)
//	}
//
//	go func() {
//		for i := 0; i < 8; i++ {
//			<-exitChan
//		}
//
//		end := time.Now().Unix()
//		fmt.Println("使用耗时: ", end-start)
//		close(primeChan)
//	}()
//
//	for {
//		_, ok := <-primeChan
//		if !ok {
//			break
//		}
//	}
//	fmt.Println("Main exit")
//}

//
//func main() {
//
//	type Cat struct {
//		Name string
//		Age  int
//	}
//	anyChan := make(chan interface{}, 3)
//
//	anyChan <- 10
//	anyChan <- "tom jack"
//	cat := Cat{"小花猫", 4}
//	anyChan <- cat
//
//	<-anyChan
//	<-anyChan
//
//	newCat := <-anyChan //从管道中取出的Cat是什么?
//	fmt.Println(newCat)
//	a := newCat.(Cat)
//	fmt.Println(a.Name)
//
//	fmt.Println(runtime.NumCPU())
//	runtime.GOMAXPROCS(2)
//	fmt.Println(runtime.NumCPU())
//
//}

//func test() {
//	//这里我们可以使用defer + recover
//	defer func() {
//		//捕获test抛出的panic
//		if err := recover(); err != nil {
//			fmt.Println("test() 发生错误", err)
//		}
//	}()
//	//定义了一个map
//	var myMap map[int]string
//	//myMap = make(map[int]string)
//	myMap[0] = "golang" //error
//}
//
//func sayHello() {
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		fmt.Println("hello,world")
//	}
//}
//
//func main() {
//	go sayHello()
//	go test()
//
//	for i := 0; i < 10; i++ {
//		fmt.Println("main() ok=", i)
//		time.Sleep(time.Second)
//	}
//
//}

//func main() {
//	var wChan chan<- int
//	wChan = make(chan int, 1)
//	wChan <- 20
//	fmt.Println(wChan)
//
//	var rChan <-chan int
//	num := <-rChan
//	fmt.Println(num)
//
//}

/*
Channel
*/
//func main() {
//
//	intChan := make(chan int, 3)
//	fmt.Println(intChan, &intChan)
//	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
//	intChan <- 10
//	intChan <- 211
//	intChan <- 50
//	<-intChan
//	intChan <- 98
//	<-intChan
//	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3
//	<-intChan
//	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3
//	<-intChan
//	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3
//	<-intChan
//}

/*
Channel Close
*/
//func main() {
//	intChan := make(chan int, 100)
//	for i := 0; i < 100; i++ {
//		intChan <- i * 2
//	}
//
//	close(intChan)
//	for v := range intChan {
//		fmt.Println("v=", v)
//	}
//}

/*
Channel Apply
*/
//func writeData(intChan chan int) {
//	for i := 0; i <= 50; i++ {
//		intChan <- i
//		fmt.Println("Write Data: ", i)
//	}
//	close(intChan)
//}
//
//func readData(intChan chan int, outChan chan bool) {
//	for {
//		v, ok := <-intChan
//		if !ok {
//			break
//		}
//		time.Sleep(time.Second)
//		fmt.Println("Read Data: ", v)
//	}
//	outChan <- true
//	close(outChan)
//}
//
//func main() {
//	intChan := make(chan int, 10)
//	outChan := make(chan bool, 1)
//
//	go writeData(intChan)
//	go readData(intChan, outChan)
//
//	for {
//		_, ok := <-outChan
//		if !ok {
//			break
//		}
//	}
//
//}

//func main() {
//	go func() {
//		for {
//			fmt.Println("hello world")
//			time.Sleep(time.Second)
//		}
//	}()
//	for i := 0; i < 11; i++ {
//		time.Sleep(time.Second)
//		fmt.Println("Hello golang ", i)
//	}
//}

//func AddSum(n int) int {
//	var sum int
//	for i := 0; i < n; i++ {
//		sum += i
//	}
//	return sum
//}

//type Monster struct {
//	Name   string  `json:"name"`
//	Age    int     `json:"age"`
//	Solary float64 `json:"solary"`
//}

//func main() {
//	var m Monster
//	var monster Monster = Monster{
//		Name:   "xiaoyang",
//		Age:    20,
//		Solary: 99.99,
//	}
//	ret, err := json.Marshal(monster)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(monster)
//	fmt.Println(string(ret))
//
//	fmt.Println("m: ", m)
//
//	_ = json.Unmarshal(ret, &m)
//	fmt.Println("m: ", m)
//}

//func main() {
//	//args := os.Args
//	//fmt.Println("------------")
//	//fmt.Println(args)
//	//fmt.Println(args[2])
//	var user string
//	var pwd string
//
//	flag.StringVar(&user, "u", "abc", "")
//	flag.StringVar(&pwd, "p", "123", "")
//	flag.Parse()
//	fmt.Println(user, pwd)
//}

//func main() {
//	fp, err := os.OpenFile("text.txt", os.O_CREATE|os.O_RDWR|os.O_RDONLY, 0644)
//
//	defer func() {
//		_ = fp.Close()
//	}()
//
//	fmt.Println(fp)
//	if err != nil {
//		panic("Open File Occur error")
//
//	}
//	//wNum, wErr := fp.WriteString("hello world")
//	//wNum1, wErr1 := fp.Write([]byte("abcd"))
//	//fmt.Println(wNum, wErr)
//	//fmt.Println(wNum1, wErr1)
//
//	reader := bufio.NewReader(fp)
//	line, _ := reader.ReadString('\n')
//	fmt.Println(line)
//
//}

//func TypeJudge(items ...interface{}) {
//	fmt.Println(items)
//	for index, v := range items {
//		switch v.(type) {
//		case int32:
//			fmt.Println("int32: ", index)
//		case float32:
//			fmt.Println("float32: ", index)
//		}
//	}
//}
//
//func main() {
//	var s1 float32 = 1.1
//	var s2 float64 = 99.99
//	var s3 int32 = 10
//	var s4 string = "hello"
//	TypeJudge(s1, s2, s3, s4)
//
//}

//type USB interface {
//	Start()
//	End()
//}
//
//type Phone struct{}
//
//type Camera struct{}
//
//type Computer struct{}
//
//func (p Phone) Start() {
//	fmt.Println("Phone start")
//}
//
//func (p Phone) End() {
//	fmt.Println("Phone end")
//}
//
//func (p Phone) Call() {
//	fmt.Println("Phone calling")
//}
//
//func (p Camera) Start() {
//	fmt.Println("Camera start")
//}
//
//func (p Camera) End() {
//	fmt.Println("Camera end")
//}
//
//func (c Computer) working(usb USB) {
//	usb.Start()
//	if phone, ok := usb.(Phone); ok {
//		phone.Call()
//	}
//	usb.End()
//}
//
//func main() {
//	var computer Computer
//	var phone Phone
//	var camera Camera
//	computer.working(phone)
//	computer.working(camera)
//
//}

//type Point struct {
//	X int
//	Y int
//}

//func main() {
//
//	var a interface{}
//	var point Point = Point{1, 2}
//	a = point
//	var b Point
//	b = a.(Point)
//	fmt.Println(b)
//	var b float32 = 1.1
//	a = b
//	x := a.(float64)
//	fmt.Println(x)
//
//	counter := make(chan int)
//	go func() {
//		counter <- 32
//
//	}()
//
//	close(counter)
//	fmt.Println(<-counter)
//}

//func fib(ch chan<- int, quit <-chan bool) {
//	x, y := 1, 1
//	for {
//		select {
//		case ch <- x:
//			x, y = y, x+y
//		case flag := <-quit:
//			fmt.Println("Flag=", flag)
//			return
//		}
//	}
//}
//
//func main() {
//	ch := make(chan int)
//	quit := make(chan bool)
//
//	go func() {
//		for i := 0; i < 8; i++ {
//			num := <-ch
//			fmt.Println("num=", num)
//		}
//		quit <- true
//	}()
//
//	fib(ch, quit)
//}
