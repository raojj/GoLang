package main

import (
	"fmt"
	"time"
)

/*
并行：parallel
并发：concurrency
Go语言内部自带并发
goroutine是go语言并行设计的核心，有人称之为go程，goroutine说到底其实就是协程，他比线程更小，十几个goroutine在底层可能就体现为
5-6个线程，go语言内部帮你实现了这些goroutine之间的内存共享
*/
func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("正在唱")
		time.Sleep(100 * time.Millisecond)
	}
}

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("正在跳")
		time.Sleep(100 * time.Millisecond)
	}
}

// main 这是主go程，main、sing、dance共同争夺cpu时间轮片，主go程结束，那子go程随之退出
func main() {
	// 两个函数并发运行了，这两个是子go程，子go程想运行，主go程不能结束
	go sing()
	go dance()
	for {
	}
}
