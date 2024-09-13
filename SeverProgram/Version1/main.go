package main

func main() {
	//创建一个sever实例
	sever := NewServer("127.0.0.1", 8888)
	sever.Start()
}
