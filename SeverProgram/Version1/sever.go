package main

import (
	"fmt"
	"net"
)

//这个最初版本创建一个服务器的类 以及开启的基本方法

// 这是一个服务器的类
type Server struct {
	Ip   string //ip地址
	Port int    //端口
}

// 创建一个接口来实现这个类的创建
func NewServer(ip string, port int) *Server {
	server := Server{
		Ip:   ip,
		Port: port,
	}
	return &server
} //传入ip和端口 接收一个服务器的类的地址

// 创建一个处理的方法
func (this *Server) Handler(conn net.Conn) {
	//当前需要处理的业务
	fmt.Println("连接建立成功")
}

// Server方法：开启服务器
func (this *Server) Start() {
	//socket 监听(下面这个是服务端的经典算法)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	//处理错误
	if err != nil { //发现错误接收到值不为空 进行如下操作
		fmt.Println("net.Listen err:", err) //返回报错结果
		return                              //终止进程
	}

	//关闭服务器
	defer listener.Close()

	//接收
	var conn net.Conn
	for { //创建一个死循环，直到错误出现
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue //跳出当前循环
		}
	}

	//处理请求
	go this.Handler(conn)

	//

}
