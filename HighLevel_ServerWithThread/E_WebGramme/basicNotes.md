
## <font face="仿宋">网络编程</font>
### <font face="仿宋">协议</font>
#### <font face="仿宋">什么是协议</font>
<font face="仿宋">一组规则，要求使用协议的双方必须严格遵守协议的内容</font>
#### <font face="仿宋">典型的协议</font>
* <font face="仿宋">传输层 常见协议有TCP/UDP协议</font>
* <font face="仿宋">应用层 常见的协议有HTTP协议、FTP协议</font>
* <font face="仿宋">网络层 常见的协议有IP协议、ICMP协议、IGMP协议</font>
* <font face="仿宋">网络接口层 常见的协议有ARP协议、RARP协议</font>
### <font face="仿宋">Socket</font>
#### <font face="仿宋">什么是Socket</font>
<font face="仿宋">Socket 英文含义是`插座、插孔`，一般称之为套接字，用于描述**IP地址和端口**。
可以实现不同程序间的数据通信。Socket也是一种`文件描述符`，具有一个类似于打开文件
的函数调用`Socket()`，该函数返回的是一个整型的Socket描述符，随后的连接建立、数据传输
等操作都是通过该Socket实现的。</font>  

### <font face="仿宋">网络应用程序设计模式</font>
#### <font face="仿宋">C/S模型</font>
<font face="仿宋">传统的网路应用设计模式，客户机（client）/服务器（server）模式。需要在通讯两端
各自部署客户机和服务器来完成数据通信。</font>
#### <font face="仿宋">B/S模式</font>
<font face="仿宋">浏览器（Browser）/服务器（Server）模式。只需要在一端部署服务器，而另一端使用
每台PC都默认设置的浏览器即可完成数据的传输</font>
#### <font face="仿宋">TCP的C/S架构</font>
|       <font face="仿宋">TCP客户端</font>       |                                             |        <font face="仿宋">TCP服务端</font>        |
| :--------------------------------------------: | :-----------------------------------------: | :----------------------------------------------: |
|                                                |                                             | <font face="Times New Roman">net.Listen()</font> |
| <font face="Times New Roman">net.Dial()</font> | <font face="仿宋">阻塞等待用户连接-></font> |   <font face="Times New Roman">Accept()</font>   |
|  <font face="Times New Roman">Write()</font>   |     <font face="仿宋"><-数据请求</font>     |    <font face="Times New Roman">Read()</font>    |
|   <font face="Times New Roman">Read()</font>   |     <font face="仿宋">数据应答-></font>     |   <font face="Times New Roman">Write()</font>    |
|  <font face="Times New Roman">Close()</font>   |                                             |   <font face="Times New Roman">Close()</font>    |

#### <font face = "仿宋">Server端</font>

```go
// Listen函数
func Listen(network, address string)(Listener, error){
    // network:选用的协议：TCP\UDP， 如tcp或者udp
    // address:IP地址+端口号， 如"127.0.0.1:8000"或":8000"
}

// Listener接口,用于监听的Socket
type Listener interface{
    Accept()(Conn, error)
    Close() error
    Addr() Addr
}

// Conn接口，用于通信的Socket
type Conn interface{
    Read(b []byte)(n int, err error)
    Write(b []byte)(n int, err error)
    Close() error
    LocalAddr() Addr
    RemoteAddr() Addr
    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}
```

#### <font face="仿宋">Server例子</font>

```go
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

}

```



#### <font face="仿宋">测试工具NetCat</font>

<font face = "仿宋">下载地址 https://eternallybored.org/misc/netcat/</font>

<font face = "仿宋">解压缩后将文件地址添加到环境变量path中</font>

<font face = "仿宋">验证：打开两个cmd窗口。</font>

<font face = "仿宋">第一个执行`nc -l -p 9000`</font>

<font face = "仿宋">第二个执行`nc localhost 9000`</font>

<font face = "仿宋">nc 127.0.0.1 8000 注意ip地址和端口之间是空格</font>

#### <font face="仿宋">Client例子</font>

```go
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

```

#### <font face = "仿宋">TCP C/S并发服务器</font>

<font face = "仿宋">net.Listen()放到for循环中，每次连接成功，创建与之对应的Socket，针对每一个Socket创建一个goroutine实现并发通信</font>
