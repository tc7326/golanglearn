package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

// Server 定义一个 服务器结构体
type Server struct {
	ip   string //ip地址
	port int    //ip端口

	//在线用户列表 一个key为string value为user的map
	onlineMap map[string]*User
	//同步锁应该是？
	mapLock sync.RWMutex

	//广播的channel
	channel chan string
}

// 类似于 Server的构造函数
func newServer(ip string, port int) *Server {

	//new 一个 Server
	server := Server{
		ip:        ip,
		port:      port,
		onlineMap: make(map[string]*User),
		channel:   make(chan string),
	}

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

	//群发处理
	go server.groupHandler()

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

	//封装 user
	user := newUser(conn, server)
	//user上线
	user.online()

	//消息处理
	go func() {
		//消息缓存
		msgBuffer := make([]byte, 2048)

		//读取消息
		for {
			n, err := conn.Read(msgBuffer)

			//没有消息了
			if n == 0 {
				user.offline()
				return
			}

			//读取IO错误
			if err != nil && err != io.EOF {
				fmt.Println("读取异常:", err)
				return
			}

			//读取到的消息 这里n前面的:一定不能少 我TM找了一下午
			msg := string(msgBuffer[:n])

			//用户消息处理
			fmt.Println("接收到客户端消息:", msg)
			user.handleMsg(msg)

		}

	}()

	//阻塞
	select {}

}

// 广播 给所有在线用户发送消息
func (server *Server) broadcast(user *User, msg string) {

	//组装消息
	readyMsg := user.name + ": " + msg

	//将消息给到 channel
	server.channel <- readyMsg

}

// 群发handler
func (server *Server) groupHandler() {

	for {

		//从服务器的channel 中取数据
		clientMsg := <-server.channel

		//遍历user 将消息交给每个user的channel
		server.mapLock.Lock()
		for _, user := range server.onlineMap {
			user.channel <- clientMsg
		}
		server.mapLock.Unlock()

	}
}
