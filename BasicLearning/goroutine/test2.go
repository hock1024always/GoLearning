package main

import (
	"fmt"
	"runtime"
	"time"
)

// 使用匿名函数加载goroutine
// 无参数状态下调用
func test1() {

	//创建一个形参为空 返回值为空的匿名函数
	go func() { //使用go 直接将这个函数作为一个goroutine来运行
		defer fmt.Println("A defer")
		func() {
			defer fmt.Println("B defer")
			runtime.Goexit() //退出这个goroutine
			//return 退出内层匿名函数
			fmt.Println("A")
		}()
	}()

	for {
		time.Sleep(1 * time.Second)
	} //死循环，使得这个test1函数无法终止
}

// 有参数状态下直接调用
func test3() {
	go func(a int, b int) {
		fmt.Println("a=", a, "b=", b)
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//test1()
	test3()
}
