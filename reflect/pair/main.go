package main

import "fmt"

// 两个接口
type Writer interface {
	WriteBook()
}
type Reader interface {
	ReadBook()
}

// 具体类型
type Book struct {
} //这个类实现了上面的两个接口

func (b *Book) WriteBook() {
	fmt.Println("写书")
}

func (b *Book) ReadBook() {
	fmt.Println("阅读")
}

func test() {
	//b_pair<type:Book ,value:>
	//b := &Book{}
}
