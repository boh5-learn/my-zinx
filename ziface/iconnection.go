package ziface

import "net"

// IConnection 定义连接博客的抽象层
type IConnection interface {
	// Start 启动连接，让当前连接准备开始工作
	Start()
	// Stop 停止连接，结束当前连接的工作
	Stop()
	// GetTCPConnection 获取当前连接的绑定 socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前连接 ID
	GetConnID() uint32
	// RemoteAddr 获取远程客户端的 TCP 状态 IP, Port
	RemoteAddr() net.Addr
	// SendMsg 发送数据
	SendMsg(msgID uint32, data []byte) error
}

// HandleFunc 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
