package main

import (
	"fmt"
)

//这个文件是来演示够浪中函数的使用的

// 测试形参和实参
func VisualValue(num int) {
	num++
}

func StaValue(num *int) {
	*num++
}
func test1() {
	fmt.Println("这是一个测试形参的函数")
	num := 10
	fmt.Println("调用函数之前的num值为：", num)
	VisualValue(num)
	fmt.Println("调用函数之后的num值为：", num)
	fmt.Println("---------------------------")

	fmt.Println("这是一个测试实参的函数")
	num1 := 10
	fmt.Println("调用函数之前的num值为：", num1)
	StaValue(&num1)
	fmt.Println("调用函数之后的num值为：", num1)
}

func main() {
	test1()
}
