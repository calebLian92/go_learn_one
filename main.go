package main

import (
	"fmt"
	"learn-one/core"
	"unsafe"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)
	fmt.Printf("init %s\n", core.IP)
	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}
	showBytes()
	testuintptr()
}
func init() {
	fmt.Println("init %s")
}

// 示例：查看字节表示
func showBytes() {
	s := "Go语言"

	// 打印每个字节的十六进制
	fmt.Printf("字符串: %s\n", s)
	fmt.Printf("字节: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%02x ", s[i])
	}
	// 输出: 47 6f e8 af ad e8 a8 80
	// G  o     语          言
	// 47=G, 6f=o, e8afad="语", e8a880="言"
}
func testuintptr() {
	// 示例1：uintptr 存储指针值
	var x int = 42
	ptr := &x // 获取x的指针

	// 将指针转换为 uintptr
	addr := uintptr(unsafe.Pointer(ptr))
	fmt.Printf("变量x的值: %d\n", x)
	fmt.Printf("变量x的地址（指针）: %p\n", ptr)
	fmt.Printf("变量x的地址（uintptr）: 0x%x\n", addr)
	fmt.Printf("uintptr大小: %d字节\n", unsafe.Sizeof(addr))
}
