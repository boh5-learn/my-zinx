package znet

import (
	"errors"
	"fmt"
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

// CallbackToClient 定义当前客户端连接所绑定的 handler(目前写死，以后优化为由用户自定义)
func CallbackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显的业务
	fmt.Println("[Conn Handle] CallbackToClient")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err:", err)
		return errors.New("CallbackToClient error")
	}

	return nil
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
		var cid uint32
		cid = 0

		// 3. 阻塞的等待客户端连接，处理业务
		for {
			// 有客户的连接，阻塞返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 将 handler 和 conn 绑定，得到我们的 Connection 模块
			dealConn := NewConnection(conn, cid, CallbackToClient)
			cid++

			// 启动当前的连接业务
			go dealConn.Start()
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
