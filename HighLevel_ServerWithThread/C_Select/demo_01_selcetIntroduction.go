package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
select:

	go里面提供了一个关键字select，通过select可以监听channel上的数据流动
	select的用法于switch语言非常类似，由select开始一个新的选择块，每个选择条件由case语句来描述。
	与switch语句相比，select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个io操作

	在一个select语句中，Go语言会按顺序从头至尾评估每一个发送和接收的语句。如果其中的任意一语句可以继续执行（即没有被阻塞）
	那么就从那些可执行的语句中任意选择一条来使用，如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有两种可能的情况：
		如果给出default语句，那么就会执行default语句，同时程序的执行会从select语句后的语句中恢复
		如果没有default语句，那么select语句将被阻塞，直到至少有一个通信可以进行下去
*/
func main() {
	ch := make(chan int)    // 用来进行数据通信的channel
	quit := make(chan bool) // 用来判断退出的channel
	// 子go程写数据
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second * 2)
		}
		close(ch)
		quit <- true      // 通知主go程退出
		runtime.Gosched() // 退出子go程
	}()

	// 主go程读数据
	for {
		select {
		case num := <-ch:
			fmt.Println("读到：", num)
		case <-quit:
			return // 中止主go程
		}
		fmt.Println("=========================")
	}
}
