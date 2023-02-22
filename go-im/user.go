package main

import (
	"fmt"
	"net"
	"strings"
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
		name:    "user_" + strings.Split(userIp, ":")[1],
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
		user.sendMsg(msg)

	}

}

// 用户上线 业务
func (user *User) online() {

	//放入 在线用户列表onlineMap
	user.server.mapLock.Lock() //加锁？
	user.server.onlineMap[user.name] = user
	user.server.mapLock.Unlock() //释放锁？

	//广播 用户上线消息
	user.server.broadcast(user, "i'm online!!!")
}

// 给用户发送消息
func (user *User) sendMsg(msg string) {
	fmt.Println("发送消息:", msg)
	user.conn.Write([]byte(msg + "\r\n"))
}

// 用户消息处理
func (user *User) handleMsg(msg string) {

	//如果用户输入 user list 则查询所有在线用户
	if msg == "user list" {

		//遍历用户列表
		user.server.mapLock.Lock()
		for _, u := range user.server.onlineMap {

			onlineMsg := u.name + " is online..."
			user.sendMsg(onlineMsg)

		}
		user.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename:" {

		//截取新用户名
		name := strings.Split(msg, ":")[1]

		//判断用户名是否重复
		_, isRepeat := user.server.onlineMap[name]
		if isRepeat {
			//提示用户 名字重复了
			user.sendMsg("repeat name!!!")
		} else {
			//改名
			user.name = name

			//改map里的值
			user.server.mapLock.Lock()
			delete(user.server.onlineMap, user.name)
			user.server.onlineMap[user.name] = user
			user.server.mapLock.Unlock()

			//回复改名成功
			user.sendMsg("rename success!!!")

		}
	} else {

		//普通消息 直接群发
		user.server.broadcast(user, msg)

	}

}

// 用户下线 业务
func (user *User) offline() {

	//从 在线用户列表onlineMap 移除
	user.server.mapLock.Lock() //加锁？
	delete(user.server.onlineMap, user.name)
	user.server.mapLock.Unlock() //释放锁？

	//广播 用户下线消息
	user.server.broadcast(user, "i'm offline!!!")

}
