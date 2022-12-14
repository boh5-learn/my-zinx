package znet

import (
	"fmt"
	"my-zinx/utils"
	"my-zinx/ziface"
	"net"
)

// NewServer 初始化 Server
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.Port,
		Router:    nil,
	}

	return s
}

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
	// Server 注册的 Router
	Router ziface.IRouter
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s. Host: %s. Port: %d\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.Port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPackageSize: %d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)

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
			dealConn := NewConnection(conn, cid, s.Router)
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

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ")
}
