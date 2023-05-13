package main

import (
	"fmt"
	"time"
)

/*
channel
是go语言中的一个核心类型，可以把它看成管道。并发核心单元通过它就可以发送或者接收数据进行通讯，这在一定程度上又进一步降低了编程的难度
channel是一个数据类型，主要用来解决协程之间数据共享（数据传递）的问题

goroutine运行在相同的地址空火箭，因此访问共享内存必须做好同步。goroutine奉行通过通信来共享内存，而不是通过共享内存来通信

引用类型channel可用于多个goroutine通讯。其内部实现了同步，确保并发安全

定义channel变量

	和map类似，channel也有一个对应的make创建的底层数据结构的引用
	当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel对象
	和其他的引用类型相似，channel的零值也是nil
	定义一个channel时，也需要定义发送到channel的值的类型。channel可以使用内置的make()函数来创建：
		chan 是创建channel所需使用的关键字。type代表指定的收发数据的类型
		make(chan Type) // 等价于make(chan Type 0)
		make(chan Type Capacity)
		当参数Capacity=0时，channel是无缓冲阻塞读写的；当Capacity>0时，channel有缓存、是非阻塞的、直到写满Capacity个元素才阻塞写入
		channel非常像生活中的管道，一边存放东西，另一边可以取东西。channel通过操作符<-来接受和发送数据

无缓冲channel，应用于两个go程之间，一个读一个写，具备同步功能
	写端写了以后。如果没有读端读取，则写端阻塞
	读端读了以后。如果没有写端写入，则读端阻塞
有缓存channel，应用于两个go程之间，一个读一个写，缓冲区可以进行数据存储，存储到容量上限，阻塞，具备异步能力
*/
// printer 定义一个打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch) // 屏幕：当有一个进程时，系统会打开这三个文件：stdin【键盘】、stdout【屏幕】、stderr
		time.Sleep(300 * time.Millisecond)
	}
}

// person1 定义两个人使用打印机
func person1() {
	printer("hello") // hello先执行
	channel <- 8     // 向channel里写入数据
}
func person2() {
	<-channel        // 从channel里面读数据但并没有保存
	printer("world") // world后执行
}

// channel 全局变量
var channel = make(chan int)

func main() {
	go person1()
	go person2()

	for {
		const n = 45
		fibN := fib(n) // slow
		fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	}
}
