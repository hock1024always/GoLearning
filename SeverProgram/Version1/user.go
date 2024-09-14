package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string //每个用户对应绑定一个自己的channl
	conn net.Conn
}

// 创建一个用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String() //用于获取网络连接的远程地址，并将其转换为字符串表示形式

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string), //创建一个新的字符串通道，用于在用户和服务器之间传递消息
		conn: conn,              //传入的参数
	}

	//启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 监听User channel的方法 一旦有消息，就发送给对端客户端（使用goroutine）
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
