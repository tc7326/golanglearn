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

}

func newUser(name string, conn net.Conn) *User {

	//user 一个 user对象
	user := User{name, conn.RemoteAddr().String(), make(chan string), conn}

	//新建一个go程 启动监听
	go func() {
		user.listenerMsg()
	}()

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