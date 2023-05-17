package main

import "fmt"

/*
死锁
	1. 单go程自己死锁
		channel至少在两个以上的go程中通信
	2. go程间channel访问顺序导致死锁
		先写后读
	3. 多go程多channel交叉死锁
	4. 在go语言中，尽量不要将互斥锁、读写锁与channel混用
*/

func main() {
	ch := make(chan int)
	//go func() {
	//	ch <- 789
	//}()
	// 写端阻塞
	ch <- 789
	num := <-ch
	fmt.Println(num)
}
