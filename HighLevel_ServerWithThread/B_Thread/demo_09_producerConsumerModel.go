package main

import (
	"fmt"
	"time"
)

/*
生产者消费者模型：写入数据--->缓冲区（公共区）--->读取数据

		成产者：
		消费者：
		缓冲区：1.解耦（降低生产者和消费者之间的耦合度，耦合度指两个板块中之间的关联程度，低耦合高内聚）
	           2.并发能力，实现多生产者（生产者和消费者之间数量不对等时，能保持正常通信）
	           3.缓存(生产者和消费者之间处理速度不同，暂存数据)
*/
func main() {
	ch := make(chan int, 2)
	// 子go程模拟生产者
	go Producer(ch)
	// 主go程模拟消费者
	Consumer(ch)
}

func Producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		fmt.Println("生产者生产：", i)
	}
	close(out)
}

func Consumer(in <-chan int) {
	for num := range in {
		fmt.Println("消费者拿到：", num)
	}
	time.Sleep(time.Second)
}
