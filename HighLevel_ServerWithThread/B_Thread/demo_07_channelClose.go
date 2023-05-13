package main

import "fmt"

/*
如果发送者知道没有更多的值需要发送到channel的话，那么让接收者也能及时知道没有多余的值可以接收将是有用的，因为接收者
可以停止不必要的的接受等待。这可以通过内置的close函数来关闭channel实现

	close(channel)
	对端可判断channel是否关闭
		if num, ok := <-channel; ok == true{}
		对端已经关闭，则返回false，num无数据
*/
func main() {
	channel := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			channel <- i
		}
		close(channel)
	}()

	for {
		if num, ok := <-channel; ok == true {
			fmt.Println("读到的数据：", num)
		} else {
			break
		}
	}
}
