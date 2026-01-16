package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个 server的接口
func NewServer(ip string, port int) *Server {
	// 相当于 new 一个server对象
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	// 当前链接的业务
	//fmt.Println("链接建立成功")

	user := NewUser(conn, this)

	// 用户上线
	user.Online()

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			// 发送0时表示下线
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// 提取用户的消息(去除'\n')
			msg := string(buf[:n-1])

			// 用户针对msg进行消息处理
			user.DoMessage(msg)
		}
	}()

	// 当前handler阻塞
	select {}
}

// 广播消息方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	// 发送消息
	this.Message <- sendMsg
}

// 监听Message广播消息channel，一旦有消息就发送给全部的user
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		// 将msg发送给全部的User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	// tcp, 127.0.0.1:8888 这个格式
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	// 方法最终结束时调用关闭
	defer listen.Close()

	// 启动监听 Message
	go this.ListenMessage()

	for {
		// accept
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err:", err)
			return
		}
		// do handler 异步处理
		go this.Handler(conn)
	}

}
