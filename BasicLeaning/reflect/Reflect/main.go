package main

import (
	"fmt"
	"reflect"
)

// 简单类型的反射
func reflectNum(arg interface{}) {
	//获取传入的变量的类型
	argType := reflect.TypeOf(arg)
	//获取传入的变量的值
	argValue := reflect.ValueOf(arg)
	fmt.Println("argType:", argType)
	fmt.Println("argValue:", argValue)
}

func test() {
	Str := "hello world"
	reflectNum(Str)
}

// 复杂类型的反射
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is called")
	fmt.Println(u.Name)
}

func test2() {
	//创建一个user类型的实例
	u := User{1, "Tom", 18}
	ReflectTest1(u)
	//
}

func ReflectTest1(input interface{}) {
	//获取input的类型
	inputType := reflect.TypeOf(input)
	fmt.Println("输入变量的类型", inputType.Name()) //这是一个TypeOf返回的type类型的方法 通过这个方法可以返回类型的名称（或种类）
	fmt.Println("输入变量的类型", inputType)        //而这个方法返回的类型是一个在全局角度的类型（也就是说会在名称前面加上所属的包）
	//获取input的值
	inputValue := reflect.ValueOf(input)
	fmt.Println("输入变量的值", inputValue)

}

func ReflectTest2(input interface{}) {
	//通过type获取里面的字段
	
}

func main() {
	//test()
	test2()
}
