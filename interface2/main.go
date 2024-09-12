package main

import "fmt"

// 下面的函数实现的是一个各种类型都可以调用的函数
func Myfunc(arg interface{}) {
	fmt.Println("我要和你")
	fmt.Println(arg)
}

func test1() {
	Myfunc("date")
	Myfunc(17)
	Myfunc(13.14)
}

func MyFunc2(arg interface{}) {
	value, ok := arg.(int) //value用来接收传入类型的值（只有断言成功才能接收到），ok用来接收类型转换的结果是否成功

	if !ok {
		fmt.Println("类型转换失败")
	} else {
		fmt.Println("类型转换成功")
		fmt.Printf("传入类型为：%T\n", value)
	}
}

func test2() {
	MyFunc2("date")
	MyFunc2(17)
	MyFunc2(13.14)
}
func main() {
	//test1()
	test2()
}
