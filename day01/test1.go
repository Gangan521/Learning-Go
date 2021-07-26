package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	// 类型转换
// 	var a int = 10
// 	var b = float32(a)
// 	fmt.Println(b)
// }

// func main() {
//终端的输入与输出
// 	var x int
// 	var y int
// 	fmt.Println("请输入一个整数，一个浮点数：")
// 	fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
// 	fmt.Printf("整数为：%d，浮点数为：%d\n", x, y)
// }

func main() {
	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读取到的数据为：", s1)
}
