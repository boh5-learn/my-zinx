package znet

import (
	"my-zinx/ziface"
	"net"
)

// Connection 连接模块
type Connection struct {
	// 当前连接的 socket
	Conn *net.TCPConn
	// 连接 ID
	ConnID uint32
	// 当前连接状态
	isClosed bool
	// 当前连接绑定的业务处理方法
	handleAPI ziface.HandleFunc
	// 告知当前连接已经退出的 channel
	ExitChan chan bool
}

// NewConnection 初始化连接
func NewConnection(conn *net.TCPConn, connID uint32, callback ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callback,
		ExitChan:  make(chan bool, 1),
	}

	return c
}
