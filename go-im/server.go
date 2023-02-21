package main

import (
	"fmt"
	"net"
)

// Server 定义一个 服务器结构体
type Server struct {
	ip   string //ip地址
	port int    //ip端口
}

// 类似于 Server的构造函数
func newServer(ip string, port int) *Server {

	//new 一个 Server
	server := Server{ip, port}

	//返回地址的值 引用传递
	return &server
}

// 定义Server的方法
func (server *Server) start() {

	//监听tcp链接
	tcp, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.ip, server.port))

	//如果发生错误
	if err != nil {
		fmt.Println("TCP监听启动失败", err)
		return
	}

	//延迟函数
	defer func(tcp net.Listener) {
		err := tcp.Close()
		if err != nil {
			fmt.Println("TCP关闭失败", err)
		}
	}(tcp)

	//处理 tcp连接
	for {

		//接受tcp连接
		conn, err := tcp.Accept()

		if err != nil {
			fmt.Println("TCP连接失败", err)
			//如果连接失败 直接跳过 下边就不处理了
			continue
		}

		//将连接交给handler处理 还要开个协程
		go server.handler(conn)

	}

}

// 处理tcp连接的逻辑
func (server *Server) handler(conn net.Conn) {
	fmt.Println("TCP连接建立成功", conn.RemoteAddr())
}
