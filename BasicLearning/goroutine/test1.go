package main

import (
	"fmt"
	"time"
)

// 创建一个方法来实现并发进程的测试
func Work1() {
	i := 0
	for {
		i++
		fmt.Println("这是第一个进程的第", i, "个循环")
		time.Sleep(1 * time.Second)
	}
}

func Work2() {
	i := 0
	for {
		i++
		fmt.Println("这是第二个进程的第", i, "个循环")
		time.Sleep(1 * time.Second)
	}
}

func Work3() {
	i := 0
	for {
		i++
		fmt.Println("这是第三个进程的第", i, "个循环")
		time.Sleep(1 * time.Second)
	}
}

//func test1() {
//	go Work1()
//	go Work2()
//	go Work3()
//	//time.Sleep(10*time.Second)
//}

func test2() {
	go Work1()

	fmt.Println("让哥们看看进程啥时候咋跑的")
	i := 0
	for {
		i++
		fmt.Println("这是主函数中进程的第", i, "个循环")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//test1()
	test2()
}
