package main

import "fmt"

//用来测试循环的程序

func test1() {
	//切片的三种表示方法
	//第一种用法 创建一个空数组当作切片来
	arr1 := []int{}
	arr1 = append(arr1, 1) //使用append方法往切片中添加一个元素

	//第二种用法 直接创建一个切片
	arr2 := []int{1, 2, 3, 4, 5}

	//第三种用法 用make函数来创建切片
	arr3 := make([]int, 5, 10) //表示一个长度为5的切片 容量为10
	//arr4 := make([]int, 5)     //表示一个长度为5的切片 容量为5

	//我们要区分两个概念 写入和追加
	arr3 = append(arr3, 100) //使用append方法往切片的末尾后一位写入一个元素
	for i := 0; i < 5; i++ {
		arr3[i] = i + 1
	} //这个是用循环将切片的已有元素修改

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//对于切片而言 访问的方法也有使用序号和值的 for range 方法 这个方法常用与未知长度的切片中
	for index, value := range arr2 {
		fmt.Println(index, value)
	}
	for _, value := range arr2 {
		fmt.Println(value)
	} //忽略前者的话需要使用_
	for index := range arr2 {
		fmt.Println(index)
	} //忽略后者的话可以省略

}

// 下面这个函数 测试使用循环对散列表的遍历
func test2() {
	//创建一个散列表
	hashmap := map[int]string{1: "a", 2: "b", 3: "c"}
	//向哈希表中添加一个键值对 直接添加即可
	hashmap[4] = "d"
	//遍历散列表
	for key, value := range hashmap {
		fmt.Println(key, value)
	}
}

// 对通道的遍历（尚未学习）
func test3() {

}

func main() {
	//test1()
	test2()
}
