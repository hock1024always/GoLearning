package main

import (
	"io"   //提供了基本的输入输出操作，例如文件读取和写入
	"log"  //提供了日志记录功能，例如打印日志信息，记录错误信息等。
	"net"  //提供了网络编程的功能，例如创建网络连接，监听端口，以及进行网络通信
	"time" //包含了时间相关的功能
)

func main() {
	listener, err := net.Listen("tcp", "IP_ADDRESS:8000") //创建一个tcp监听器，监听IP_ADDRESS的8000端口
	if err != nil {
		log.Fatal(err)
	} //上面的err变量接受的是Listen函数返回的任意错误信息，如果err不为空，则表示出现了错误，使用log.Fatal函数打印错误信息并终止程序的执行。

	for {
		conn, err := listener.Accept() //监听端口并接受连接，返回一个新的连接对象和错误信息
		if err != nil {
			log.Print(err)
			continue
		} //如果err不为空，则表示出现了错误，使用log.Print函数打印错误信息并继续下一次循环，等待下一个连接。

		handleConn(conn) //启动一个新的goroutine处理连接，将连接对象传递给handleConn函数。
	}
}

func handleConn(c net.Conn) {
	defer c.Close() //确保在函数结束时关闭连接，释放资源。

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) //将当前时间写入连接对象c中，使用io.WriteString函数将时间字符串写入连接。
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second) //暂停1秒钟，等待下一次循环。
	}
}
