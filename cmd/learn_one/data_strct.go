package main

import (
	"fmt"
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
func main() {
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
	multiSlicePassing()
}
