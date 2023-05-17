package main

import (
	"fmt"
	"sync"
	"time"
)

/*
互斥锁: 同一时刻只能有一个go程访问数据资源

建议锁：操作系统提供，建议你在编程时使用
强制锁：编程用不到，系统自己用的
*/

// 使用channel完成同步
/*func printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
}

func person1() {
	printer("hello")
	ch <- 789
}

func person2() {
	<-ch
	printer("world")
}

var ch = make(chan int)

func main() {
	go person1()
	go person2()
	for {
		;
	}
}*/
func printer(str string) {
	mutex.Lock()
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

// 创建一个互斥锁，新建的互斥锁状态为0，未加锁
var mutex sync.Mutex

func main() {
	go person1()
	go person2()
	for {

	}
}
