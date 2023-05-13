package main

import (
	"fmt"
	"time"
)

/*
周期定时：
type Ticker struct{
	C<-chan Time
	r runTimeTimer
}
*/

func main() {
	quit := make(chan bool)
	myTicker := time.NewTicker(time.Second)
	fmt.Println("now", time.Now())
	i := 0
	go func() {
		for {
			nowTime := <-myTicker.C
			i++
			fmt.Println("nowTime:", nowTime)
			if i == 8 {
				quit <- true
			}
		}
	}()
	<-quit
}
