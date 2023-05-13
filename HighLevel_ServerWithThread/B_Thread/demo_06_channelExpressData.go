package main

/*
len(channel):返回channel剩余未读取的内容的长度
cap(channel):返回channel的容量
*/
import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("i=", i)
		}
		// 通知主go程打印结束
		channel <- "打印结束"
	}()
	for {
		if "打印结束" == <-channel {
			// fmt.Println(<-channel) // fatal error: all goroutines are asleep - deadlock!for循环上已经读取了，所以此时构成了死锁
			break
		}
	}
}
