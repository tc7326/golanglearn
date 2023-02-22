package main

import (
	"fmt"
	"net"
)

// User 定义一个 用户结构体
type User struct {
	name    string      //用户名
	addr    string      //网络地址
	channel chan string //消息通道
	conn    net.Conn    //tcp连接

	server *Server //对应的服务器

}

func newUser(conn net.Conn, server *Server) *User {

	//获取客户端ip
	userIp := conn.RemoteAddr().String()

	//user 一个 user对象
	user := User{
		name:    "user-" + userIp,
		addr:    userIp,
		channel: make(chan string),
		conn:    conn,
		server:  server}

	//新建一个go程 启动监听
	go user.listenerMsg()

	return &user

}

func (user *User) listenerMsg() {

	for {
		//从chan里取消息
		msg := <-user.channel

		//将去到的消息重新回复过去
		user.conn.Write([]byte(fmt.Sprint(user.name, msg)))

	}

}

// 用户上线 业务
func (user *User) online() {

	//放入 在线用户列表onlineMap
	user.server.mapLock.Lock() //加锁？
	user.server.onlineMap[user.name] = user
	user.server.mapLock.Unlock() //释放锁？

	//广播 用户上线消息
	user.server.broadcast(user, " is online!!!\n")
}

// 用户群聊 业务
func (user *User) sendMsg(msg string) {

	user.server.broadcast(user, msg+"\n")

}

// 用户下线 业务
func (user *User) offline() {

	//从 在线用户列表onlineMap 移除
	user.server.mapLock.Lock() //加锁？
	delete(user.server.onlineMap, user.name)
	user.server.mapLock.Unlock() //释放锁？

	//广播 用户下线消息
	user.server.broadcast(user, " is offline!!!\n")

}
