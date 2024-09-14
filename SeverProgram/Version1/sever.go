package main

import (
	"fmt"
	"net"
	"sync"
)

//这个最初版本创建一个服务器的类 以及开启的基本方法

// 这是一个服务器的类
type Server struct {
	Ip   string //ip地址
	Port int    //端口

	//在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex //与 同步 相关的读写锁一般都在sync的包里面

	//消息广播的channel
	Message chan string
}

// 创建一个接口来实现这个类的创建
func NewServer(ip string, port int) *Server {
	server := Server{
		Ip:   ip,
		Port: port,

		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return &server
} //传入ip和端口 接收一个服务器的类的地址

// 监听message广播消息的goroutine和channel，一旦有消息出现就推送给所有的在线user
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		//将msg发送给全部的在线User的channel中
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {

	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 创建一个处理的方法
func (this *Server) Handler(conn net.Conn) {
	//当前需要处理的业务
	//fmt.Println("连接建立成功")

	user := NewUser(conn)
	//用户上线，将用户加入OnlineMap中
	this.mapLock.Lock()
	this.OnlineMap[user.Addr] = user
	this.mapLock.Unlock()

	//广播用户上线的消息
	this.BroadCast(user, "上线了")

	//当前channel阻塞
	select {}

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

	//关闭监听的socket
	defer listener.Close()

	//启动监听message的goroutine
	go this.ListenMessage()

	//接收 Conn是一个网络连接 net.Conn 接口定义了一组方法，用于在网络上读写数据、设置读写超时时间、以及获取连接的本地和远程地址信息等
	for { //创建一个死循环，直到错误出现
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue //跳出当前循环
		} //监听不到消息或者在接受连接时发生错误，err 的值将是非 nil 的，它会包含具体的错误信息

		//处理请求
		go this.Handler(conn)

	}

}
