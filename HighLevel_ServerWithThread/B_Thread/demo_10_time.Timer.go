package main

import (
	"fmt"
	"time"
)

/*
time.Timer:
Timer是一个定时器。代表未来的一个单一时间，你可以告诉timer你要等待多长时间
type Timer struct{
	C <-chan Time
	r runtimeTimer
}
提供一个channel，在定时时间到达之前，没有数据写入,timer.C会一直阻塞。直到定时时间到，系统会自动向time.C这个channel中写入当前时间，阻塞
即被解除

定时器的停止和重置：
myTimer.Stop()停止定时器，但是没有结束定时器
myTimer.Reset(time.Second)

*/

/*
三种定时的方法：
1. time.Sleep(time.Second)
2. myTimer := time.NewTimer(time.Second * 2)
3. time.After(time.Second)
*/
func main() {
	// 获取当前系统时间
	fmt.Println("系统当前时间：", time.Now())
	// 创建一个定时器
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <-myTimer.C
	fmt.Println("系统现下时间：", nowTime)
	nowTime1 := <-time.After(time.Second)
	fmt.Println("系统现下时间：", nowTime1)
}
