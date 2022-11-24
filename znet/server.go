package znet

import (
	"fmt"
	"io"
	"my-zinx/ziface"
	"net"
)

// Server 实现 IServer, Server 服务器模块
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的 ip 版本
	IPVersion string
	// 服务器监听的 IP
	IP string
	// 服务器监听的端口
	Port int
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	go func() {
		// 1. 获取 TCP 的 Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}

		// 2. 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp error:", err)
			return
		}

		fmt.Printf("start zinx server succ, %s succ, listenning\n", s.Name)

		// 3. 阻塞的等待客户端连接，处理业务
		for {
			// 有客户的连接，阻塞返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 已经与客户端建立连接，做一些业务，做一个最大 512 字节的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err == io.EOF {
						fmt.Println("conn closed")
						break
					}
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}

					fmt.Printf("recv client buf: %s, cnt: %d\n", buf[:cnt], cnt)

					// 回显功能
					if _, err := conn.Write(buf[0:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	// TODO 将一些服务器资源、状态或者开辟的链接信息进行停止或回收
}

func (s *Server) Serve() {
	// 启动 Server 服务
	s.Start()

	// TODO 做一些启动服务器后的额外业务

	// 阻塞状态
	select {}
}

// NewServer 初始化 Server
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}
