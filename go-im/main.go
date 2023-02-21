package main

func main() {

	//开启一个服务端
	server := newServer("localhost", 25567)

	//启动服务端
	server.start()

}
