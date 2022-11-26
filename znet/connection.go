package znet

import (
	"fmt"
	"io"
	"my-zinx/utils"
	"my-zinx/ziface"
	"net"
)

// NewConnection 初始化连接
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}

	return c
}

// Connection 连接模块
type Connection struct {
	// 当前连接的 socket
	Conn *net.TCPConn
	// 连接 ID
	ConnID uint32
	// 当前连接状态
	isClosed bool
	// 告知当前连接已经退出的 channel
	ExitChan chan bool
	// 该连接的 Router
	Router ziface.IRouter
}

// StartReader 读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running")
	defer fmt.Printf("ConnID: %d, reader is exit, remote addr is: %s\n", c.ConnID, c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端数据到 buf 中
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		cnt, err := c.Conn.Read(buf)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("recv buf err:", err)
			continue
		}

		// 得到 Request
		req := Request{
			conn: c,
			data: buf[:cnt],
		}

		// 从 Router 中调用业务
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	fmt.Printf("Conn start. ConnID: %d\n", c.ConnID)

	// 读数据
	go c.StartReader()

	// TODO 启动从当前连接写数据的业务
}

func (c *Connection) Stop() {
	fmt.Printf("Conn stop. ConnID: %d\n", c.ConnID)

	// 如果当前连接已关闭
	if c.isClosed {
		return
	}

	c.isClosed = true

	// 关闭 socket 连接
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("close conn err:", err)
	}

	// 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	// TODO implement me
	panic("implement me")
}
