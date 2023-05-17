package main

import (
	"fmt"
	"net"
)

func HandlerConnect(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		// 注意这里读入的数据里会有一个\n
		if "exit\n" == string(buf[:n]) {
			fmt.Println("服务器接收到客户端的断开连接请求...")
			return
		}
		if n == 0 {
			fmt.Println("服务器检测到客户端已经关闭，断开连接")
			return
		}
		if err != nil {
			fmt.Println("conn.Read error:", err)
			return
		}
		addr := conn.RemoteAddr()
		fmt.Printf("服务器读到来自%s的数据:%s\n：", addr, string(buf[:n]))
		// 给客户发送数据
		_, err = conn.Write([]byte("hello this is Server"))
		if err != nil {
			fmt.Println("conn.Write error:", err)
			return
		}
	}
}

func main() {
	// 创建监听Socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)

	// 监听客户端连接请求
	for {
		fmt.Println("服务器等待客户端连接")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		// 具体完成服务器和客户端的通信
		go HandlerConnect(conn)
	}
}
