package main

import "fmt"

// channel实现了goroutine实现两个进程之间的通信
//chan1 := make(chan int)     //make(chan Type)无缓冲通道
//chan2 := make(chan int, 10) //make(chan Type, capacity)有缓冲通道

//向管道中写入数据
//channel <- value 发送

//从管道中读取数据
//<- channel 接收并将其丢弃
//x := <- channel 接收并赋值
//x , ok := <- channel 接收并赋值，ok为false表示channel已关闭

func test1() {
	defer fmt.Println("主进程已经结束")
	//创建一个无缓冲channel
	c := make(chan int)

	go func() {
		defer fmt.Println("子进程已经结束")
		fmt.Println("正在进行")
		num := 666
		fmt.Println("子进程中数据的值为：", num)
		c <- num
	}()

	num := <-c //通过通道将子进程的数据捕捉到主进程
	fmt.Println("主进程捕捉到的子进程数据： ", num)
}

func main() {
	test1()
}
