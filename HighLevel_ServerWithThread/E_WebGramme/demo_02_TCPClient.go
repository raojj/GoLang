package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器IP+Port 创建通信Socket
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	// 给服务器发送数据
	_, err = conn.Write([]byte("hello this is Client"))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return
	}
	buf := make([]byte, 4099)
	// 接收服务器回发的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}
	fmt.Println("这是服务器发过来的数据", string(buf[:n]))
}
