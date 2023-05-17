package main

import (
	"fmt"
	"net"
)

func main() {
	// 127.0.0.1本地回环地址
	// 指定服务器的通讯协议和IP地址、port，返回用于监听的Socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			return
		}
	}(listener)
	fmt.Println("服务器等待客户端建立连接...")
	// 阻塞监听用户端连接请求，成功建立连接，返回用于通信的Socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept error:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	fmt.Println("服务器与客户端成功建立连接...")
	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}
	fmt.Println("服务器读到数据：", string(buf[:n]))

	// 给客户发送数据
	_, err = conn.Write([]byte("hello this is Server"))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return
	}
}
