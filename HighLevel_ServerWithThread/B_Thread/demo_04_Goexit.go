package main

/*
return后，后面的函数语句不会再执行
Goexit 退出当前go程
GOMAXPROCS 用来设置当前进程可以使用的最大cpu核数
*/
import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		// defer延迟调用，延迟到函数结束的时候
		defer fmt.Println("aaaaaaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbb")
	}()
	for {

	}
}
func test() {
	defer fmt.Println("ccccccccccccccccc")
	// 匿名函数中的test函数没有加go关键词，test里的Goexit会把匿名函数的go程关掉
	runtime.Goexit()
	fmt.Println("dddddddddddddddddddd")
}
