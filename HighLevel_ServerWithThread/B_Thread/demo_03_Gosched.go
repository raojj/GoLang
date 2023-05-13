package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Gosched():出让当前go程所占用的cpu时间片,cpu调度器会调用其他等待运行的任务运行，并在下次再获得cpu时间片的时候继续运行
*/

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("-----------this is goroutine test------------")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		// 让出cpu时间片
		runtime.Gosched()
		fmt.Println("-----------this is main Test------------")
		time.Sleep(100 * time.Millisecond)
	}
}
