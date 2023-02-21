package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello4() {

	//goRunShow()

	showChannel()

	showChannel10()
}

// goroutine
func goRunTask() {
	i := 0
	for {
		i++
		fmt.Println(time.Now(), "new goroutine", i)
		defer fmt.Println(time.Now(), "new goroutine defer", i)
		time.Sleep(500 * time.Millisecond)
		if i == 5 {
			fmt.Println(time.Now(), "new is end", i)
			defer fmt.Println(time.Now(), "new goroutine defer", i)
			runtime.Goexit()
		}
	}
}

func goRunShow() {

	//go程一个循环
	go goRunTask()

	//主线程一个循环
	i := 0
	for {
		i++
		fmt.Println(time.Now(), "main goroutine", i)
		time.Sleep(500 * time.Millisecond)
		if i == 10 {
			fmt.Println(time.Now(), "main is end", i)
			return
		}
	}

}

// channel 应该就是一个缓存队列的概念 用于协程见共享数据
func showChannel() {

	//没缓存的
	channel0 := make(chan string)

	go func() {
		defer fmt.Println("子go程结束")
		fmt.Println("子go程正在运行……")

		//发送数据到 channel
		channel0 <- "哈哈"
	}()

	//取数据
	str := <-channel0
	fmt.Println("str = ", str)
	fmt.Println("main go程结束")

}

func showChannel10() {

	//有缓存的
	channel10 := make(chan int, 10)
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(channel10), cap(channel10))

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 3; i++ {
			channel10 <- i //往channel里放数据
			fmt.Printf("子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(channel10), cap(channel10))
		}
	}()

	time.Sleep(2 * time.Second)

	//遍历数据
	for i := 0; i < 3; i++ {
		num := <-channel10 //从channel中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main进程结束")

}
