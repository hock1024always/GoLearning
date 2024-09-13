package main

import "fmt"

// 从下面这个案例我们可以发现 defer语句在触发defer关键字的时候将语句压入栈 之后在函数结束的时候（或者遇到return语句时开始出栈）
func test() {

	fmt.Println("Start")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	fmt.Println("Middle")

	defer func() {
		fmt.Println("Deferred with function")
	}()

	fmt.Println("End")
}

func test2() {

	fmt.Println("Start")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	return

	fmt.Println("Middle")

	defer func() {
		fmt.Println("Deferred with function")
	}()

	fmt.Println("End")
}

func main() {
	//test2()
}
