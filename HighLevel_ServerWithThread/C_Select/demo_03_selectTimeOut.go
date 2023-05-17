package main

import (
	"fmt"
	"time"
)

/*
使用select解决超时问题
有时候会出现goroutine阻塞的情况，那么我们可以利用select来设置超时
*/

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("time out")
				o <- true
			}
		}
	}()
	c <- 666
	<-o
}
