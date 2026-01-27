package main

import (
	"fmt"
	"learn-one/cmd/server"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

/*
(a) 指针类型（Pointer）

(b) 数组类型

(c) 结构化类型(struct)

(d) Channel 类型

(e) 函数类型

(f) 切片类型

(g) 接口类型（interface）

(h) Map 类型
举例说明一下。
*/
//
//var a = "阿里云大学"
//var b string = "edu.aliyun.com"
//var c bool

func main1() {

	//structExample()
	//channelExample()
	//functionExample()
	//sliceExample()
	//interfaceExample()
	//mapExample()
	//fmt.Println(a, b, c)
}
func pointerExample() {
	fmt.Println("=== 指针类型 ===")

	var x int = 10
	var y string = "hello"

	// 声明指针
	var p1 *int    // 指向int的指针
	var p2 *string // 指向string的指针

	p1 = &x // 获取x的地址
	p2 = &y

	fmt.Printf("x的值: %d, 地址: %p\n", x, p1)
	fmt.Printf("y的值: %s, 地址: %p\n", y, p2)

	// 通过指针修改值
	*p1 = 20
	*p2 = "world"

	fmt.Println("修改后:", x, y) // 20 world

	// 空指针
	var p3 *int
	if p3 == nil {
		fmt.Println("p3是空指针")
	}
}
func arrayExample() {
	fmt.Println("\n=== 数组类型 ===")

	// 声明数组（固定长度）
	var arr1 [3]int              // 默认值 [0 0 0]
	var arr2 = [3]int{1, 2, 3}   // 初始化
	arr3 := [...]int{4, 5, 6, 7} // 自动推断长度

	fmt.Println("arr1:", arr1) // [0 0 0]
	fmt.Println("arr2:", arr2) // [1 2 3]
	fmt.Println("arr3:", arr3) // [4 5 6 7]

	// 访问和修改
	arr1[0] = 10
	arr1[1] = 20
	fmt.Println("修改后arr1:", arr1)

	// 多维数组
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("二维数组:", matrix)
	fmt.Println("matrix[1][2]:", matrix[1][2]) // 6

	// 数组是值类型（复制时会复制整个数组）
	arr4 := arr2
	arr4[0] = 100
	fmt.Println("arr2:", arr2) // [1 2 3] - 不变
	fmt.Println("arr4:", arr4) // [100 2 3]
}
func structExample() {
	fmt.Println("\n=== 结构体类型 ===")

	// 定义结构体
	type Person struct {
		Name    string
		Age     int
		Email   string
		Married bool
	}

	// 创建结构体实例
	var p1 Person // 零值
	p1.Name = "Alice"
	p1.Age = 25
	p1.Email = "alice@example.com"

	// 字面量初始化
	p2 := Person{
		Name:    "Bob",
		Age:     30,
		Email:   "bob@example.com",
		Married: true,
	}

	// 简短初始化（按字段顺序）
	p3 := Person{"Charlie", 35, "charlie@example.com", false}

	fmt.Printf("p1: %+v\n", p1)
	fmt.Printf("p2: %+v\n", p2)
	fmt.Printf("p3: %+v\n", p3)

	// 匿名结构体
	point := struct {
		X, Y int
		Name string
	}{
		X:    10,
		Y:    20,
		Name: "坐标点",
	}
	fmt.Println("匿名结构体:", point)

	// 嵌套结构体
	type Address struct {
		City, Street, ZipCode string
	}

	type Employee struct {
		Person // 匿名字段（嵌入）
		Address
		Salary float64
		ID     int
	}

	emp := Employee{
		Person: Person{
			Name: "David",
			Age:  28,
		},
		Address: Address{
			City:    "北京",
			Street:  "中关村",
			ZipCode: "100080",
		},
		Salary: 8000.50,
		ID:     1001,
	}

	fmt.Printf("\n员工信息:\n")
	fmt.Printf("姓名: %s (可直接访问: %s)\n", emp.Person.Name, emp.Name)
	fmt.Printf("地址: %s, %s\n", emp.City, emp.Street)
	fmt.Printf("薪资: %.2f\n", emp.Salary)
}

func channelExample() {
	fmt.Println("\n=== Channel类型 ===")

	// 创建channel
	ch1 := make(chan int)       // 无缓冲channel
	ch2 := make(chan string, 3) // 有缓冲channel，容量3

	// 使用goroutine发送数据
	go func() {
		fmt.Println("开始发送数据...")
		ch1 <- 100
		ch1 <- 200
		close(ch1) // 关闭channel
	}()

	// 主goroutine接收数据
	fmt.Println("接收数据:")
	for i := range ch1 {
		fmt.Println("从ch1接收到:", i)
	}

	// 有缓冲channel示例
	ch2 <- "Hello"
	ch2 <- "World"
	ch2 <- "!"
	close(ch2)

	fmt.Println("\n从缓冲channel接收:")
	for msg := range ch2 {
		fmt.Println("从ch2接收到:", msg)
	}

	// Channel方向
	//var sendOnly chan<- int = make(chan<- int, 1)  // 只写
	//var receiveOnly <-chan int                     // 只读

	// select多路复用
	ch3 := make(chan int)
	ch4 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch3 <- 1
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch4 <- "fast"
	}()

	select {
	case v := <-ch3:
		fmt.Println("从ch3接收到:", v)
	case v := <-ch4:
		fmt.Println("从ch4接收到:", v) // 先执行这个
	case <-time.After(200 * time.Millisecond):
		fmt.Println("超时")
	}

}

func functionExample() {
	fmt.Println("\n=== 函数类型 ===")

	// 函数类型声明
	type MathFunc func(int, int) int
	type StringProcessor func(string) string

	// 实现函数
	add := func(a, b int) int {
		return a + b
	}

	multiply := func(a, b int) int {
		return a * b
	}

	// 使用函数类型变量
	var operation MathFunc
	operation = add
	fmt.Println("10 + 5 =", operation(10, 5))
	fmt.Println("10 + 5 =", add(10, 5))

	operation = multiply
	fmt.Println("10 * 5 =", operation(10, 5))

	// 函数作为参数
	calculate := func(fn MathFunc, x, y int) int {
		return fn(x, y)
	}

	fmt.Println("计算 10+5:", calculate(add, 10, 5))
	fmt.Println("计算 10*5:", calculate(multiply, 10, 5))

	// 函数作为返回值
	getOperator := func(op string) MathFunc {
		switch op {
		case "+":
			return add
		case "*":
			return multiply
		default:
			return func(a, b int) int { return 0 }
		}
	}

	op := getOperator("+")
	fmt.Println("使用返回的函数:", op(20, 30)) // 50

	// 闭包示例
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("\n闭包计数器:")
	fmt.Println("计数:", counter()) // 1
	fmt.Println("计数:", counter()) // 2
	fmt.Println("计数:", counter()) // 3
}
func sliceExample() {
	fmt.Println("\n=== 切片类型 ===")
	//cap 容量
	// 创建切片的不同方式
	var s1 []int               // nil切片
	s2 := []int{1, 2, 3, 4, 5} // 字面量
	s3 := make([]int, 3)       // 长度3，容量3
	s4 := make([]int, 3, 5)    // 长度3，容量5

	fmt.Printf("s1: %v, len=%d, cap=%d, nil? %v\n", s1, len(s1), cap(s1), s1 == nil)
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4: %v, len=%d, cap=%d\n", s4, len(s4), cap(s4))

	// 切片操作
	fmt.Println("\n切片操作:")
	s5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("s5[2:5]:", s5[2:5]) // [2 3 4]
	fmt.Println("s5[:3]:", s5[:3])   // [0 1 2]
	fmt.Println("s5[7:]:", s5[7:])   // [7 8 9]
	fmt.Println("s5[:]:", s5[:])     // 全部

	// 追加元素
	s6 := []int{1, 2, 3}
	s6 = append(s6, 4, 5, 6)
	fmt.Println("\n追加后 s6:", s6) // [1 2 3 4 5 6]

	// 切片复制
	src := []int{10, 20, 30, 40, 50}
	dst := make([]int, 3)
	copy(dst, src)               // 只复制前3个
	fmt.Println("复制后 dst:", dst) // [10 20 30]

	// 二维切片
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("\n二维切片:")
	for i, row := range matrix {
		fmt.Printf("行%d: %v\n", i, row)
	}

	// 切片与数组的关系
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]               // 切片引用数组的部分
	slice[0] = 99                   // 修改切片会影响原数组
	fmt.Println("\n数组 arr:", arr)   // [1 99 3 4 5]
	fmt.Println("切片 slice:", slice) // [99 3 4]
}

// 定义接口
type Speaker interface {
	Speak() string
}

type Animal interface {
	Speaker
	Move() string
}

// 实现接口的结构体
type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// Dog 的方法
func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

func (d Dog) Move() string {
	return d.Name + " is running"
}

// Cat 的方法
func (c Cat) Speak() string {
	return c.Name + " says: Meow!"
}

func (c Cat) Move() string {
	return c.Name + " is jumping"
}

func interfaceExample() {
	fmt.Println("\n=== 接口类型 ===")

	// 使用接口
	var s Speaker
	var a Animal

	dog := Dog{"Buddy"}
	cat := Cat{"Kitty"}

	s = dog
	fmt.Println(s.Speak()) // Buddy says: Woof!

	a = cat
	fmt.Println(a.Speak()) // Kitty says: Meow!
	fmt.Println(a.Move())  // Kitty is jumping

	// 空接口（可以存储任何类型）
	var empty interface{}
	empty = 42
	fmt.Println("\n空接口存储int:", empty)

	empty = "hello"
	fmt.Println("空接口存储string:", empty)

	empty = []int{1, 2, 3}
	fmt.Println("空接口存储切片:", empty)

	// 类型断言
	var value interface{} = 3.14

	if v, ok := value.(float64); ok {
		fmt.Printf("是float64: %f\n", v)
	}

	// 类型switch
	switch v := value.(type) {
	case int:
		fmt.Println("是int:", v)
	case float64:
		fmt.Println("是float64:", v) // 执行这个
	case string:
		fmt.Println("是string:", v)
	default:
		fmt.Println("未知类型")
	}
}
func mapExample() {
	fmt.Println("\n=== Map类型 ===")

	// 创建map的不同方式
	var m1 map[string]int      // nil map
	m2 := make(map[string]int) // 空map
	m3 := map[string]int{      // 字面量初始化
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}

	fmt.Println("m1 (nil):", m1)
	fmt.Println("m2 (空):", m2)
	fmt.Println("m3 (有数据):", m3)

	// 操作map
	m3["grape"] = 4 // 添加或修改
	fmt.Println("添加后 m3:", m3)

	// 获取值
	appleCount := m3["apple"]
	fmt.Println("apple数量:", appleCount) // 5

	// 检查key是否存在
	if count, exists := m3["watermelon"]; exists {
		fmt.Println("watermelon数量:", count)
	} else {
		fmt.Println("watermelon不存在")
	}

	// 检查key是否存在
	if count, exists := m3["apple"]; exists {
		fmt.Println("apple数量:", count)
	} else {
		fmt.Println("apple不存在")
	}

	// 删除元素
	delete(m3, "banana")
	fmt.Println("删除banana后:", m3)

	// 遍历map
	fmt.Println("\n遍历map:")
	for key, value := range m3 {
		fmt.Printf("  %s: %d\n", key, value)
	}

	// 复杂value的map
	type Person struct {
		Age  int
		City string
	}

	people := map[string]Person{
		"Alice":   {Age: 25, City: "北京"},
		"Bob":     {Age: 30, City: "上海"},
		"Charlie": {Age: 28, City: "广州"},
	}

	fmt.Println("\n人员信息:")
	for name, info := range people {
		fmt.Printf("  %s: %d岁, 来自%s\n", name, info.Age, info.City)
	}

	// map的map
	scores := map[string]map[string]int{
		"math": {
			"Alice": 95,
			"Bob":   88,
		},
		"english": {
			"Alice": 92,
			"Bob":   85,
		},
	}

	tt := 1
	fmt.Println("\nmap的map:", tt)
	var a int = 10
	var b = 10
	c := 10
	fmt.Println("a:", a, "b:", b, "c:", c)
	fmt.Println("\n学生成绩:")
	for subject, studentScores := range scores {
		fmt.Printf("  %s:\n", subject)
		for student, score := range studentScores {
			fmt.Printf("    %s: %d\n", student, score)
		}
	}
}

var x, y int

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	a int
	b bool
)
var c, d int = 1, 2
var e, f = 123, "hello"

func main2() {
	//这种不带声明格式的只能在函数体中出现//g, h := 123, "hello"
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
func valueType() {
	/*// 基本数据类型
	bool
	int, int8, int16, int32, int64
	uint, uint8, uint16, uint32, uint64
	float32, float64
	complex64, complex128
	string  // 注意：string是值类型，但有特殊行为
	byte    // uint8的别名
	rune    // int32的别名

	// 复合类型
	array      // 数组
	struct     // 结构体*/
	// 基本类型的值传递
	a := 10
	b := a // 复制值
	b = 20
	fmt.Println(a, b) // 输出: 10 20，a不变

	// 数组的值传递
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // 复制整个数组
	arr2[0] = 100
	fmt.Println(arr1) // [1 2 3]
	fmt.Println(arr2) // [100 2 3]

	// 结构体的值传递
	type Point struct {
		X, Y int
	}
	p1 := Point{10, 20}
	p2 := p1 // 复制整个结构体
	p2.X = 100
	fmt.Println(p1) // {10 20}
	fmt.Println(p2) // {100 20}

	var f string = "abc"
	//a := "20" //不允许
	f = "20" //允许，因为这是给相同的变量赋予一个新的值
	fmt.Println("hello, world", f)
}

const (
	a1 = "abc"
	b1 = len(a1)
	c1 = unsafe.Sizeof(a1)
)

func nonBlocking() {
	ch := make(chan string, 1)

	// 非阻塞发送
	select {
	case ch <- "消息":
		fmt.Println("发送成功")
	default:
		fmt.Println("通道已满，发送失败")
	}

	// 非阻塞接收
	select {
	case msg := <-ch:
		fmt.Println("收到:", msg)
	default:
		fmt.Println("没有消息")
	}
}

func multipleChannels() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		messages := []string{"A", "B", "C", "D", "E"}
		for _, msg := range messages {
			ch2 <- msg
			time.Sleep(300 * time.Millisecond)
		}
		close(ch2)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	for {
		select {
		case num, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 关闭")
				ch1 = nil // 设置为 nil 后，这个 case 将不再被选中
			} else {
				fmt.Printf("从 ch1 收到数字: %d\n", num)
			}

		case str, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 关闭")
				ch2 = nil
			} else {
				fmt.Printf("从 ch2 收到字符串: %s\n", str)
			}

		case <-done:
			fmt.Println("收到结束信号")
			return
		}
	}
}
func loopWithSelect() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("定时器触发:", t.Format("15:04:05"))
		case <-done:
			fmt.Println("完成")
			ticker.Stop()
			return
		}
	}
}

func testFallthrough() {
	score := 85

	fmt.Printf("分数 %d: ", score)
	switch {
	case score >= 90:
		fmt.Print("优秀")
		fallthrough // 继续执行下一个 case
	case score >= 80:
		fmt.Print("，良好")
		fallthrough
	case score >= 70:
		fmt.Print("，中等")
		fallthrough
	case score >= 60:
		fmt.Print("，及格")
	default:
		fmt.Print("，不及格")
	}
	// 输出：分数 85: ，良好，中等，及格
}
func swap(x, y string) (string, string) {
	return y, x
}

func newInitialization() {
	// new() 返回指针，并初始化为零值
	p1 := new(int)     // *int, 指向 0
	p2 := new(string)  // *string, 指向 ""
	p3 := new([3]int)  // *[3]int, 指向 [0, 0, 0]
	p4 := new(struct { // 指向结构体零值
		x int
		y string
	})

	*p1 = 100
	fmt.Println(*p1, p2, *p3, *p4)
}
func makeInitialization() {
	// 1. 切片初始化
	slice1 := make([]int, 5)     // 长度=5, 容量=5: [0,0,0,0,0]
	slice2 := make([]int, 3, 10) // 长度=3, 容量=10: [0,0,0]

	// 2. 映射初始化
	map1 := make(map[string]int)     // 空map
	map2 := make(map[string]int, 10) // 指定初始容量

	// 3. 通道初始化
	ch1 := make(chan int)        // 无缓冲通道
	ch2 := make(chan string, 10) // 缓冲大小为10

	fmt.Println("切片:", slice1, slice2)
	fmt.Println("映射:", map1, map2)
	fmt.Println("通道:", ch1, ch2)
}

// 复合字面量初始化
func compositeLiterals() {
	// 1. 结构体初始化
	type Point struct {
		X, Y int
	}
	p1 := Point{10, 20}       // 必须按顺序
	p2 := Point{X: 10, Y: 20} // 指定字段名
	p3 := Point{X: 10}        // Y 使用零值: 0
	p4 := Point{}             // X:0, Y:0

	// 2. 数组初始化
	arr1 := [3]int{1, 2, 3}      // [1, 2, 3]
	arr2 := [3]int{1: 20, 2: 30} // [0, 20, 30]
	arr3 := [...]int{1, 2, 3, 4} // 编译器计算长度

	// 3. 切片初始化
	slice1 := []int{1, 2, 3}
	slice2 := []string{"a", "b", "c"}

	// 4. 映射初始化
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[int]string{1: "one", 2: "two"}

	// 5. 接口存储具体类型
	var iface interface{} = "hello"

	fmt.Println("结构体:", p1, p2, p3, p4)
	fmt.Println("数组:", arr1, arr2, arr3)
	fmt.Println("切片:", slice1, slice2)
	fmt.Println("映射:", map1, map2)
	fmt.Println("接口:", iface)
}

func trapMultipleReturn() {
	// 正确：同时声明多个变量
	a, b := 1, 2

	// 错误：不能重复声明
	//a, b := 3, 4 // 编译错误：a 已声明

	// 正确：只要有一个新变量就可以
	// 至少有一个新变量就可以
	a, c := 3, 4 // ✅ a 被重新赋值，c 是新变量

	// 使用 _ 忽略某些值
	//_, d := someFunction()

	fmt.Println(a, b, c, d)
}
func bestPractices() {
	// 1. 尽量使用简短声明
	name := "Alice" // ✅ 好
	var age = 25    // ✅ 好
	var height int  // ⚠️ 只在需要零值时使用

	// 2. 接近使用位置声明变量
	// ✅ 好：在需要使用时声明
	for i := 0; i < 10; i++ {
		result := i * 2
		fmt.Println(result)
	}

	// 3. 初始化时赋予有意义的值
	var count int = 0       // ✅ 明确初始值
	var total = calculate() // ✅ 立即初始化

	// 4. 对于复杂类型，使用 make 预分配
	// 如果知道大概大小，预分配可以提高性能
	data := make([]int, 0, 100) // 长度0，容量100
	cache := make(map[string]int, 1000)

	// 5. 使用零值的优势
	var users []User // nil 切片，可以安全 append

	fmt.Println(users, data, cache, name, age, height, total, count)
}

func calculate() int {
	return 42
}

type User struct {
	Name string
}

// =================================
// 重要：Go 中数组是值类型！
// 传递给函数时会发生复制
// =================================

func basicArrayPassing() {
	// 定义一个数组
	nums := [5]int{1, 2, 3, 4, 5}

	fmt.Println("调用前:", nums)

	// 传递数组给函数（会发生复制）
	modifyArray(nums)

	fmt.Println("调用后:", nums) // 仍然是 {1, 2, 3, 4, 5}
}

func modifyArray(arr [5]int) {
	fmt.Println("函数内修改前:", arr)
	arr[0] = 100 // 修改的是副本
	arr[2] = 300
	fmt.Println("函数内修改后:", arr)
}

// =================================
// 传递多维数组
// =================================
func multiArrayPassing() {
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("调用前 matrix:")
	printMatrix(matrix)

	// 传递二维数组（值传递，会复制）
	modifyMatrix(matrix)

	fmt.Println("调用后 matrix:")
	printMatrix(matrix) // 原数组不变
}

func printMatrix(mat [2][3]int) {
	for _, row := range mat {
		fmt.Println(row)
	}
}

func modifyMatrix(mat [2][3]int) {
	fmt.Println("函数内修改 matrix:")
	mat[0][0] = 100
	mat[1][2] = 600
	printMatrix(mat)
}

// 指针传递
func arrayPointerPassing() {
	// =================================
	// 方法1：传递数组指针
	// =================================
	nums := [5]int{1, 2, 3, 4, 5}

	fmt.Println("调用前:", nums)

	// 传递数组的指针
	modifyArrayByPointer(&nums)

	fmt.Println("调用后:", nums) // 修改生效了！
}

func modifyArrayByPointer(arr *[5]int) {
	// 通过指针修改原数组
	(*arr)[0] = 100 // 等价于 arr[0] = 100（Go的语法糖）
	arr[2] = 300    // Go 自动解引用

	fmt.Println("函数内:", *arr)
}

// =================================
// 传递多维数组指针
// =================================
func multiArrayPointerPassing() {
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("调用前 matrix:")
	printMatrix(matrix)

	// 传递二维数组指针
	modifyMatrixByPointer(&matrix)

	fmt.Println("调用后 matrix:")
	printMatrix(matrix) // 修改生效了！
}

func modifyMatrixByPointer(mat *[2][3]int) {
	// 修改原数组
	(*mat)[0][0] = 100
	mat[1][2] = 600 // 自动解引用

	fmt.Println("函数内修改后的 matrix:")
	printMatrix(*mat)
}

// 切片传递
func slicePassing() {
	// =================================
	// 推荐：使用切片而不是数组
	// 切片是引用类型，传递时不会复制底层数据
	// =================================

	// 创建数组
	arr := [5]int{1, 2, 3, 4, 5}

	// 转换为切片
	slice := arr[:] // 引用整个数组

	fmt.Println("调用前 slice:", slice)

	// 传递切片
	modifySlice(slice)

	fmt.Println("调用后 slice:", slice) // 修改生效！
	fmt.Println("原数组 arr:", arr)     // 数组也被修改了！
}

func modifySlice(s []int) {
	s[0] = 100
	s[2] = 300

	// 注意：可以追加元素
	// s = append(s, 6)  // 这会创建新切片，需要返回

	fmt.Println("函数内 slice:", s)
}

// =================================
// 传递多维切片
// =================================
func multiSlicePassing() {
	// 创建二维切片
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("调用前 matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// 传递二维切片
	modifyMultiSlice(matrix)

	fmt.Println("调用后 matrix:")
	for _, row := range matrix {
		fmt.Println(row) // 修改生效！
	}
}

func modifyMultiSlice(mat [][]int) {
	// 修改元素
	mat[0][0] = 100
	mat[1][2] = 600

	// 甚至可以修改行
	mat[2] = []int{700, 800, 900}

	fmt.Println("函数内修改后的 matrix:")
	for _, row := range mat {
		fmt.Println(row)
	}
}
func variableLengthArrays() {
	// =================================
	// Go 中数组长度是类型的一部分
	// 不同长度的数组是不同的类型
	// 使用切片来处理可变长度
	// =================================

	// 错误：不能将不同长度的数组传递给同一个函数
	arr3 := [3]int{1, 2, 3}
	arr5 := [5]int{1, 2, 3, 4, 5}
	//printSlice(arr3)  // 需要 printArray3
	//printSlice(arr5)  // 需要 printArray5

	printArray3(arr3) // 需要 printArray3
	printArray5(arr5) // 需要 printArray5

	// 正确：使用切片
	slice3 := []int{1, 2, 3}
	slice5 := []int{1, 2, 3, 4, 5}

	printSlice(slice3) // 可以处理任意长度
	printSlice(slice5)
}

// 可以处理任意长度的切片
func printSlice(s []int) {
	fmt.Println("切片长度:", len(s), "容量:", cap(s))
	for i, v := range s {
		fmt.Printf("s[%d] = %d\n", i, v)
	}
}

// 如果必须使用数组，需要为每个长度写一个函数
func printArray3(arr [3]int) {
	fmt.Println(arr)
}

func printArray5(arr [5]int) {
	fmt.Println(arr)
}

// /性能考虑和最佳实践
func performanceConsideration() {
	// =================================
	// 性能测试：数组 vs 切片传递
	// =================================

	// 大数组（1MB）
	var bigArray [1000000]int // 约 8MB

	// 测试值传递（数组）
	start := time.Now()
	passByValue(bigArray)
	fmt.Printf("值传递耗时: %v\n", time.Since(start))

	// 测试指针传递
	start = time.Now()
	passByPointer(&bigArray)
	fmt.Printf("指针传递耗时: %v\n", time.Since(start))

	// 测试切片传递
	bigSlice := bigArray[:]
	start = time.Now()
	passBySlice(bigSlice)
	fmt.Printf("切片传递耗时: %v\n", time.Since(start))
}

func passByValue(arr [1000000]int) {
	// 这里会复制整个数组（8MB）
	_ = arr[0]
}

func passByPointer(arr *[1000000]int) {
	// 只传递指针（8字节）
	_ = arr[0]
}

func passBySlice(s []int) {
	// 切片头部（24字节）
	_ = s[0]
}

// =================================
// 最佳实践建议
// =================================
func bestPractices2() {
	// 建议1：小数组可以直接传递
	smallArray := [3]int{1, 2, 3}
	processSmallArray(smallArray) // 复制开销小，没问题

	// 建议2：大数组使用指针或切片
	largeArray := [10000]int{}

	// 方式A：传递指针（明确表示会修改）
	processLargeArrayByPointer(&largeArray)

	// 方式B：传递切片（更灵活）
	processLargeArrayBySlice(largeArray[:])

	// 建议3：多维数组使用切片
	// 而不是 [N][M]int，使用 [][]int
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	processMatrix(matrix)
}

func processSmallArray(arr [3]int) {
	// 处理小数组
}

func processLargeArrayByPointer(arr *[10000]int) {
	// 通过指针修改
}

func processLargeArrayBySlice(s []int) {
	// 通过切片处理
}

func processMatrix(mat [][]int) {
	// 处理二维切片
}

func main5() {
	//println(a1, b1, c1)
	//var a int = 4
	//var b int32
	//var c float32
	//var ptr *int
	//
	///* 运算符实例 */
	//fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	//fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	//fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)
	//
	///*  & 和 * 运算符实例 */
	//ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	//fmt.Printf("a 的值为  %d\n", a)
	//fmt.Printf("*ptr 为 %d\n", *ptr)

	//nonBlocking()
	//multipleChannels()
	//loopWithSelect()
	//testFallthrough()

	//a, b := swap("Mahesh", "Kumar")
	//fmt.Println(a, b)
	//newInitialization()
	//makeInitialization()
	//compositeLiterals()
	//trapMultipleReturn()
	//bestPractices()
	//basicArrayPassing()
	//multiArrayPassing()
	//arrayPointerPassing()
	//slicePassing()
	//multiSlicePassing()
	//variableLengthArrays()
	//bestPractices2()
	performanceConsideration()
}

// 指针
func main6() {
	//var a int = 20 /* 声明实际变量 */
	//var ip *int    /* 声明指针变量 */
	//
	//ip = &a /* 指针变量的存储地址 */
	//
	//fmt.Printf("a 变量的地址是: %x\n", &a)
	//
	///* 指针变量的存储地址 */
	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	//
	///* 使用指针访问值 */
	//fmt.Printf("*ip 变量的值: %d\n", *ip)

	//pointArray()
	//pointerArrayOperations()
	//structPointerArray()
	//dynamicPointerArray()
	functionExamples()
}

// 指针数组与函数
// 修改指针数组中的值
func modifyPointerArray(arr [3]*int, multiplier int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			*arr[i] *= multiplier
		}
	}
}

// 返回指针数组
func createPointerArray(values []int) []*int {
	ptrs := make([]*int, len(values))
	for i := range values {
		ptrs[i] = &values[i]
	}
	return ptrs
}

func functionExamples() {
	fmt.Println("\n=== 5. 指针数组与函数 ===")

	// 函数修改指针数组
	vals := [3]int{5, 10, 15}
	var ptrArr [3]*int
	for i := range vals {
		ptrArr[i] = &vals[i]
	}

	fmt.Println("修改前:", *ptrArr[0], *ptrArr[1], *ptrArr[2])
	modifyPointerArray(ptrArr, 2)
	fmt.Println("修改后:", *ptrArr[0], *ptrArr[1], *ptrArr[2])

	// 函数返回指针数组
	newValues := []int{50, 100, 150, 200}
	newPtrArr := createPointerArray(newValues)

	fmt.Println("\n函数返回的指针数组:")
	for i, ptr := range newPtrArr {
		fmt.Printf("  [%d]: %d\n", i, *ptr)
	}
}

// 动态数组与指针
func dynamicPointerArray() {
	fmt.Println("\n=== 4. 动态数组与指针 ===")

	// 使用切片创建动态指针数组
	size := 3
	dynamicPtrs := make([]*int, size)

	// 分配值
	values := []int{100, 200, 300}
	for i := 0; i < size; i++ {
		dynamicPtrs[i] = &values[i]
	}

	fmt.Println("动态指针数组:")
	for i, ptr := range dynamicPtrs {
		fmt.Printf("  索引 %d: 值=%d\n", i, *ptr)
	}

	// 追加新元素
	newValue := 400
	dynamicPtrs = append(dynamicPtrs, &newValue)

	fmt.Println("\n追加后:")
	for i, ptr := range dynamicPtrs {
		fmt.Printf("  索引 %d: 值=%d\n", i, *ptr)
	}
}
func pointArray() {
	fmt.Println("=== 1. 指针数组基本声明 ===")

	// 普通数组
	var arr = [5]int{10, 20, 30, 40, 50}

	// 指针数组 - 存储int类型指针的数组
	var ptrArr [5]*int

	// 将数组中每个元素的地址赋给指针数组
	for i := 0; i < len(arr); i++ {
		ptrArr[i] = &arr[i]
	}

	fmt.Printf("原始数组: %v\n", arr)
	fmt.Printf("指针数组内容: %v\n", ptrArr)

	// 通过指针访问值
	for i := 0; i < len(ptrArr); i++ {
		fmt.Printf("ptrArr[%d] = %p, *ptrArr[%d] = %d\n",
			i, ptrArr[i], i, *ptrArr[i])
	}
}

// 指针数组的初始化与操作
func pointerArrayOperations() {
	fmt.Println("\n=== 2. 指针数组的初始化与操作 ===")

	// 方法1: 声明后逐个赋值
	var ptrArr1 [3]*int
	a, b, c := 1, 2, 3
	ptrArr1[0] = &a
	ptrArr1[1] = &b
	ptrArr1[2] = &c

	// 方法2: 直接初始化
	x, y, z := 10, 20, 30
	ptrArr2 := [3]*int{&x, &y, &z}

	fmt.Println("ptrArr1 指向的值:", *ptrArr1[0], *ptrArr1[1], *ptrArr1[2])
	fmt.Println("ptrArr2 指向的值:", *ptrArr2[0], *ptrArr2[1], *ptrArr2[2])

	// 通过指针修改值
	*ptrArr1[0] = 100
	fmt.Println("修改后 a 的值:", a)

	// 指针数组的遍历
	fmt.Println("\n遍历 ptrArr2:")
	for index, ptr := range ptrArr2 {
		fmt.Printf("索引 %d: 指针=%p, 值=%d\n", index, ptr, *ptr)
	}
}

// 结构体指针数组
type Person struct {
	Name string
	Age  int
}

func structPointerArray() {
	fmt.Println("\n=== 3. 结构体指针数组 ===")

	// 创建结构体指针数组
	people := [3]*Person{
		&Person{"Alice", 25},
		&Person{"Bob", 30},
		&Person{"Charlie", 35},
	}

	// 访问和修改
	fmt.Println("初始人员信息:")
	for i, p := range people {
		fmt.Printf("  %d: %s, %d岁\n", i, p.Name, p.Age)
		//  fmt.Println((*p).Name, (*p).Age)     // 显式解引用 ✅
	}

	// 修改结构体内容
	people[0].Age = 26
	people[1] = &Person{"David", 28}

	fmt.Println("\n修改后的人员信息:")
	for i, p := range people {
		fmt.Printf("  %d: %s, %d岁\n", i, p.Name, p.Age)
	}

	// 检查nil指针
	var personArr [2]*Person
	personArr[0] = &Person{"Eve", 22}
	// personArr[1] 是 nil

	fmt.Println("\n检查nil指针:")
	for i, p := range personArr {
		if p != nil {
			fmt.Printf("  personArr[%d]: %s\n", i, p.Name)
		} else {
			fmt.Printf("  personArr[%d]: nil指针\n", i)
		}
	}
}

// 结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main7() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "edu.aliyun.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "edu.aliyun.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)
}
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

type Books1 struct {
	title   string
	author  string
	subject string
	book_id int
}

func main8() {
	var Book1 Books1 /* Declare Book1 of type Book */
	var Book2 Books1 /* Declare Book2 of type Book */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "edu.aliyun.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "edu.aliyun.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook1(&Book1)

	/* 打印 Book2 信息 */
	printBook1(&Book2)
}
func printBook1(book *Books1) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book title : %s\n", (*book).title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
func main9() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice2(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice2(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice2(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice2(number3)
}
func printSlice2(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main11() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

// range
func main12() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
func main13() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	/* 使用 key 输出 map 值 */
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}
}
func main14() {
	/* 创建 map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}

// 斐波那契
func fibonaci(n int) int {
	if n < 2 {
		return n
	}
	return fibonaci(n-2) + fibonaci(n-1)
}
func main15() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonaci(i))
	}
}
func main17() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %.2f\n", mean)
	//referenceAndInterfaceTypes()
	//explicitConversion()
	//stringConversion()
	//typeAssertion()
	//customTypeConversion()
	//pointerConversion()
	mapSliceConversionLimits()
}
func mapSliceConversionLimits() {
	fmt.Println("\n=== 7. map 和 slice 的转换限制 ===")

	// 7.1 切片不能直接转换类型
	intSlice := []int{1, 2, 3}
	// floatSlice := []float64(intSlice)  // 错误

	// 需要逐个元素转换
	floatSlice := make([]float64, len(intSlice))
	for i, v := range intSlice {
		floatSlice[i] = float64(v)
	}
	fmt.Printf("[]int%v → []float64%v\n", intSlice, floatSlice)

	// 7.2 map 键值类型转换
	m1 := map[string]int{"a": 1, "b": 2}
	// m2 := map[interface{}]int(m1)  // 错误

	// 需要创建新的map
	m2 := make(map[interface{}]int)
	for k, v := range m1 {
		m2[k] = v
	}
	fmt.Printf("map[string]int%v → map[interface{}]int%v\n", m1, m2)
}
func pointerConversion() {
	fmt.Println("\n=== 5. 指针类型转换 ===")

	// 5.1 unsafe.Pointer 转换
	var i int = 42
	//var f float64 = 3.14

	// 普通指针不能直接转换类型
	// var p *float64 = (*float64)(&i)  // 编译错误

	pi := &i
	pf := (*float64)(unsafe.Pointer(pi))
	fmt.Printf("int指针 %p → float64指针 %p\n", pi, pf)
	fmt.Printf("危险！原值 %d，转换后值 %f\n", i, *pf)

	// 5.2 不同结构体指针转换
	type A struct{ x int }
	type B struct{ x int }

	a := A{x: 42}
	pb := (*B)(unsafe.Pointer(&a))
	fmt.Printf("A{x:%d} → B{x:%d}\n", a.x, pb.x)
}

// 断言
func typeAssertion() {
	fmt.Println("\n=== 3. 接口类型断言 ===")

	// 语法: value, ok := interfaceValue.(Type)

	var any interface{} = "Hello, Go!"

	// 3.1 安全类型断言
	if str, ok := any.(string); ok {
		fmt.Printf("断言成功: 是字符串, 值: %s\n", str)
	} else {
		fmt.Println("断言失败: 不是字符串")
	}

	// 3.2 另一种写法（不安全，可能panic）
	any = 42
	// num := any.(int)  // 如果any不是int，会panic

	// 3.3 处理多种类型
	processInterface(42)
	processInterface("Hello")
	processInterface(3.14)

	// 3.4 类型switch
	fmt.Println("\n类型switch:")
	var values []interface{} = []interface{}{42, "hello", 3.14, true}
	for _, v := range values {
		switch val := v.(type) {
		case int:
			fmt.Printf("int: %d\n", val)
		case string:
			fmt.Printf("string: %s\n", val)
		case float64:
			fmt.Printf("float64: %.2f\n", val)
		default:
			fmt.Printf("未知类型: %T\n", val)
		}
	}
}

func processInterface(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("处理整数: %d × 2 = %d\n", v, v*2)
	case string:
		fmt.Printf("处理字符串: 长度 = %d\n", len(v))
	default:
		fmt.Printf("不支持的类型: %T\n", v)
	}
}

// 字符串转化
func stringConversion() {
	fmt.Println("\n=== 2. 字符串转换 ===")

	// 2.1 字符串 ↔ 字节切片
	str := "Hello, 世界"

	// 字符串转字节切片
	bytes := []byte(str)
	fmt.Printf("字符串 → 字节切片: %v\n", bytes)
	fmt.Printf("长度: 字符串=%d, 字节切片=%d\n", len(str), len(bytes))

	// 字节切片转字符串
	str2 := string(bytes)
	fmt.Printf("字节切片 → 字符串: %s\n", str2)

	// 2.2 字符串 ↔ 符文切片
	runes := []rune(str)
	fmt.Printf("\n字符串 → 符文切片: %v\n", runes)
	fmt.Printf("符文数: %d (中文字符占1个符文)\n", len(runes))

	str3 := string(runes)
	fmt.Printf("符文切片 → 字符串: %s\n", str3)

	// 2.3 数值 ↔ 字符串（使用strconv包）
	num := 42
	numStr := strconv.Itoa(num) // int to string
	fmt.Printf("\nint(%d) → 字符串: %s (类型: %T)\n", num, numStr, numStr)

	// 字符串转数值
	strNum := "123"
	parsedNum, err := strconv.Atoi(strNum)
	if err == nil {
		fmt.Printf("字符串(\"%s\") → int: %d\n", strNum, parsedNum)
	}

	// 浮点数转换
	pi := 3.14159
	piStr := strconv.FormatFloat(pi, 'f', 2, 64)
	fmt.Printf("float64(%.5f) → 字符串: %s\n", pi, piStr)

	pi2, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("字符串(\"3.14\") → float64: %.2f\n", pi2)
}

// 显示转化
func explicitConversion() {
	fmt.Println("\n=== 1. 显式类型转换 ===")

	// 语法: T(expression)

	// 1.1 数值类型之间的转换
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("int(%d) → float64: %.2f\n", i, f)
	fmt.Printf("float64(%.2f) → uint: %d\n", f, u)

	// 1.2 整数和浮点数转换
	var x float64 = 3.14
	var y int = int(x) // 截断小数部分
	fmt.Printf("float64(%.2f) → int: %d (截断)\n", x, y)

	// 1.3 不同大小的整数转换
	var a int16 = 255
	var b int32 = int32(a)
	var c int8 = int8(a) // 可能溢出
	fmt.Printf("int16(%d) → int32: %d\n", a, b)
	fmt.Printf("int16(%d) → int8: %d (溢出!)\n", a, c)

	// 1.4 字节和整数
	var ch byte = 'A'
	var ascii int = int(ch)
	fmt.Printf("byte('%c') → int: %d\n", ch, ascii)

	// 1.5 符文和整数
	var r rune = '中'
	var codePoint int = int(r)
	fmt.Printf("rune('%c') → int: %d (Unicode码点)\n", r, codePoint)
}
func referenceAndInterfaceTypes() {
	fmt.Println("\n=== 引用类型和接口类型 ===")

	// 3.1 引用类型（切片、映射、通道、函数、指针）
	//slice := []int{1, 2, 3}    // 引用类型
	//m := map[string]int{}      // 引用类型
	//ch := make(chan int)       // 引用类型

	// 3.2 指针类型
	var x int = 10
	var ptr *int = &x
	//var ptrToSlice *[]int = &slice

	// 3.3 函数类型
	var fn func(int, int) int
	fn = func(a, b int) int {
		return a + b
	}

	// 3.4 接口类型
	var any interface{} = "可以是任何类型"
	//var reader interface {
	//	Read([]byte) (int, error)
	//}

	// 3.5 类型别名和自定义类型
	type MyInt int     // 自定义类型，新类型
	type YourInt = int // 类型别名，同一种类型

	var mi MyInt = 42
	var yi YourInt = 42

	fmt.Printf("指针: %T, 指向的值: %d\n", ptr, *ptr)
	fmt.Printf("函数: %T\n", fn)
	fmt.Printf("接口: %T, 值: %v\n", any, any)
	fmt.Printf("自定义类型: %T, 类型别名: %T\n", mi, yi)

	// 检查类型
	fmt.Println("\n类型检查:")
	fmt.Printf("MyInt 底层类型: %v\n", reflect.TypeOf(mi).Kind())
	fmt.Printf("YourInt 底层类型: %v\n", reflect.TypeOf(yi).Kind())
}

func customTypeConversion() {
	fmt.Println("\n=== 4. 自定义类型转换 ===")

	// 4.1 类型别名 - 可以直接转换
	type Celsius float64
	type Fahrenheit float64

	var c Celsius = 100
	var f Fahrenheit = Fahrenheit(c*9/5 + 32)

	fmt.Printf("%.2f°C = %.2f°F\n", c, f)

	// 4.2 底层类型相同但需要显式转换
	type Meter float64
	type Kilometer float64

	var m Meter = 1000
	var km Kilometer = Kilometer(m / 1000)
	fmt.Printf("%.2f 米 = %.2f 千米\n", m, km)

	// 4.3 结构体类型转换（通常不行，除非字段完全相同）
	type Point1 struct{ X, Y int }
	type Point2 struct{ X, Y int }

	p1 := Point1{10, 20}
	// p2 := Point2(p1)  // 错误：不能直接转换

	// 需要显式转换每个字段
	p2 := Point2{p1.X, p1.Y}
	fmt.Printf("Point1%v → Point2%v\n", p1, p2)
}

// 接口
type Phone interface {
	call()
}
type NokiaPhone struct{}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct{}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func main18() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

// main.go - 完整版本 支付

// 定义接口（所有包都能看到）
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type OrderManager interface {
	UpdateOrderStatus(orderID string, status string) error
	OrderPaid(orderID string) error
}

// OrderServiceAdapter 适配器，让 order.OrderService 实现 OrderManager 接口
type OrderServiceAdapter struct {
	orderService *server.OrderService
}

func (a *OrderServiceAdapter) UpdateOrderStatus(orderID string, status string) error {
	return a.orderService.UpdateOrderStatus(orderID, status)
}

func (a *OrderServiceAdapter) OrderPaid(orderID string) error {
	return a.orderService.OrderPaid(orderID)
}

// PaymentServiceAdapter 适配器，让 payment.PaymentService 实现 PaymentProcessor 接口
type PaymentServiceAdapter struct {
	paymentService *server.PaymentService
}

func (a *PaymentServiceAdapter) ProcessPayment(amount float64) error {
	return a.paymentService.ProcessPayment(amount)
}

func main19() {
	fmt.Println("=== 开始演示双向依赖解决方案 ===")

	// 1. 创建基础服务实例
	baseOrderService := server.NewOrderService()
	basePaymentService := server.NewPaymentService()

	// 2. 创建适配器  orderService: baseOrderService    类型 ： 值
	orderAdapter := &OrderServiceAdapter{orderService: baseOrderService}
	paymentAdapter := &PaymentServiceAdapter{paymentService: basePaymentService}

	// 3. 设置双向依赖
	// 订单服务需要支付处理器
	baseOrderService.SetPaymentProcessor(paymentAdapter)
	//baseOrderService.SetPaymentProcessor(basePaymentService)
	// 支付服务需要订单管理器
	basePaymentService.SetOrderManager(orderAdapter)

	fmt.Println("\n=== 测试创建订单 ===")
	// 4. 使用服务
	err := baseOrderService.CreateOrder(100.0)
	if err != nil {
		fmt.Printf("创建订单失败: %v\n", err)
	} else {
		fmt.Println("创建订单成功!")
	}

	fmt.Println("\n=== 测试直接调用支付 ===")
	err = basePaymentService.ProcessPayment(50.0)
	if err != nil {
		fmt.Printf("支付失败: %v\n", err)
	} else {
		fmt.Println("支付成功!")
	}
}

// 不适用适配器
// 订单服务接口
type OrderService interface {
	CreateOrder() error
	UpdateOrderStatus(orderID string, status string) error
}

// 支付服务接口
type PaymentService interface {
	ProcessPayment(amount float64) error
}

// OrderServiceImpl 实现
type OrderServiceImpl struct {
	paymentService PaymentService
}

func (o *OrderServiceImpl) CreateOrder() error {
	fmt.Println("创建订单")
	return o.paymentService.ProcessPayment(100.0)
}

func (o *OrderServiceImpl) UpdateOrderStatus(orderID string, status string) error {
	fmt.Printf("更新订单状态: %s -> %s\n", orderID, status)
	return nil
}

func (o *OrderServiceImpl) SetPaymentService(ps PaymentService) {
	o.paymentService = ps
}

// PaymentServiceImpl 实现
type PaymentServiceImpl struct {
	orderService OrderService
}

func (p *PaymentServiceImpl) ProcessPayment(amount float64) error {
	fmt.Printf("处理支付: %.2f\n", amount)
	return p.orderService.UpdateOrderStatus("order_123", "paid")
}

func (p *PaymentServiceImpl) SetOrderService(os OrderService) {
	p.orderService = os
}

func main20() {
	// 创建实例
	orderSvc := &OrderServiceImpl{}
	paymentSvc := &PaymentServiceImpl{}

	// 关键：相互注入
	orderSvc.SetPaymentService(paymentSvc) // PaymentServiceImpl 实现了 PaymentService
	paymentSvc.SetOrderService(orderSvc)   // OrderServiceImpl 实现了 OrderService

	// 使用
	orderSvc.CreateOrder()
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 	`error` 接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
} // 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
func main() {

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
