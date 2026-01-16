package main

/**
编译命令：
go build -o server.exe main.go server.go user.go
*/

func main() {
	// 启动服务
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
