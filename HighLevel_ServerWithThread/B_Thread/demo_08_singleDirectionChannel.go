package main

import "fmt"

/*
默认的channel是双向的 make(chan type)
单向写channel	var sendCh chan<- Type
单向读channel	var revCh <-chan Type
转换：

	双向的channel可以转化为任意一种单向channel
		sendCh = ch
	单向channel不能转化为双向channel

传参：

	传【引用】
*/
/*
func main() {
	ch := make(chan int)       // 双向channel
	var sendCh chan<- int = ch // 单向写channel
	sendCh <- 789 // deadLock 只有一个goroutine
	// num := <-sendCh  "Invalid operation: <-sendCh (receive from the send-only type chan<- int)"
	var revCh <-chan int = ch // 单向读channel
	<-revCh // deadLock 只有一个goroutine

	// 反向赋值
	// var ch2 chan int = sendCh  "Cannot use 'sendCh' (type chan<- int) as the type chan int"
}
*/
func send(out chan<- int) {
	out <- 89
	close(out)
}

func recv(in <-chan int) {
	num := <-in
	fmt.Println(num)
}
func main() {
	ch := make(chan int)

	go func() {
		send(ch) // 双向channel转化为单向写channel
	}()
	recv(ch)
}
