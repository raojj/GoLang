package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
条件变量：

	本身不是锁，但是可以和锁一起使用

go标准库中的sync.Cond类型代表了条件变量。条件变量要与锁(互斥锁，或者读写锁)一起使用。成员变量L代表与条件变量搭配使用的锁

	type Cond struct {
		noCopy noCopy
		// L is held while observing or changing the condition
		L Locker
		notify  notifyList
		checker copyChecker
	}

对应的三种常用方法：Wait、Signal、Broadcast
func(c *Cond) Wait()

	该函数的作用可以归纳为三点：（调用wait函数前一定要先加锁）
		阻塞等待条件变量曼珠
		释放已掌握的互斥锁，相当于cond.L.Unlock()，注意：两步为一个原子操作
		当被唤醒，Wait()函数返回时，解除阻塞并重新获取互斥锁。相当于cond.L.Lock()

func(c *Cond) Signal()（唤醒一个）

	单发通知：给一个正在等待（阻塞）在该条件变量上的goroutine发送通知

func(c *Cond) Broadcast()（唤醒所有）

	广发通知：给正在等待（阻塞）在该条件变量上的所有goroutine发送通知
*/
/*
使用流程：
1. 创建条件变量：var cond sync.Cond
2. 指定条件变量用的锁： cond.L = new(sync.Mutex)
3. cond.L.Lock() 给公共区域加锁
4. 判断是否达到阻塞条件（缓冲区满或空）----for循环判断
5. 访问公共区域----读写数据、打印
6. 公共区域解锁
7. 唤醒等待在条件变量上的goroutine
*/
var cond sync.Cond // 创建全局条件变量

func producer(out chan<- int, idx int) {
	for {
		cond.L.Lock()       // 条件变量对应互斥锁加锁
		for len(out) == 3 { // 产品区满，等待消费者消费
			cond.Wait() // 挂起当前进程，等待条件变量满足，被消费者唤醒
		}
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%dth生产者，产生数据%3d，公共区域剩余%d个数据\n", idx, num, len(out))
		cond.L.Unlock()         // 生产结束，解锁互斥锁
		cond.Signal()           // 唤醒阻塞的消费者
		time.Sleep(time.Second) // 生产完休息一下，给别的goroutine执行机会
	}
}

func consumer(in <-chan int, idx int) {
	for {
		cond.L.Lock()      // 条件变量对应互斥锁加锁
		for len(in) == 0 { // 产品区空，等待生产者生产
			cond.Wait() // 挂起当前进程，等待条件变量满足，被生产者唤醒
		}
		num := <-in
		fmt.Printf("%dth消费者，消费数据%3d，公共区域剩余%d个数据\n", idx, num, len(in))
		cond.L.Unlock()         // 生产结束，解锁互斥锁
		cond.Signal()           // 唤醒阻塞的消费者
		time.Sleep(time.Second) // 生产完休息一下，给别的goroutine执行机会
	}
}

func main() {
	product := make(chan int, 5)
	cond.L = new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go consumer(product, i)
	}
	for i := 0; i < 5; i++ {
		go producer(product, i)
	}
	for {

	}
}
