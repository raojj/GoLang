package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
读写锁：
	读写锁可以让多个操作并发，同时读取，但是对于写操作是完全互斥的，也就是说，当一个goroutine进行写操作的时候，其他goroutine既不能进行读操作，也不能进行写操作
	go中的读写锁由结构体类型sync.RWMutex表示。此类型的方法集合中包含两对方法：
		一组是对写操作的锁定和解锁，简称"写锁定"和"写解锁"
		func (*RWMutex)Lock()
		func (*RWMutex)Unlock()
		另一组表示对读操作的锁定和解锁，简称为"读锁定"与"读解锁"
		func (*RWMutex)RLock()
		func (*RWMutex)RUlock()
*/

func readGo06(in <-chan int, idx int) {
	for {
		num := <-in
		fmt.Printf("---%dth读go程，读出：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
	}
}
func writeGo06(out chan<- int, idx int) {
	for {
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%dth写go程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
	}
}
func main() {
	ch1 := make(chan int, 5)
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	// 用于数据传递
	// 启动5个读的go程
	for i := 0; i < 5; i++ {
		go writeGo06(ch1, i+1)
	}
	// 启动5个写的go程
	for i := 0; i < 5; i++ {
		go readGo06(ch1, i+1)
	}
	for {

	}
}
