package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("[Client] start...")

	time.Sleep(1 * time.Second)

	// 1. 连接 server，得到 conn
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	// 2. 调用 Write 写数据
	for {
		_, err := conn.Write([]byte("Hello Zinx v0.1"))
		if err != nil {
			fmt.Println("write conn err:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
			return
		}

		fmt.Printf("server call back: %s, cnt: %d\n", buf[:cnt], cnt)

		// 阻塞
		time.Sleep(1 * time.Second)
	}
}
