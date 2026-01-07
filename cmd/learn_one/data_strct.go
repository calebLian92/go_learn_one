package main

import (
	"fmt"
	"time"
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
func main() {
	//structExample()
	//channelExample()
	//functionExample()
	//sliceExample()
	//interfaceExample()
	mapExample()
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

	fmt.Println("\n学生成绩:")
	for subject, studentScores := range scores {
		fmt.Printf("  %s:\n", subject)
		for student, score := range studentScores {
			fmt.Printf("    %s: %d\n", student, score)
		}
	}
}
