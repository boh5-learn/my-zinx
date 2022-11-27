package znet

import (
	"errors"
	"fmt"
	"io"
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
		// buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		// cnt, err := c.Conn.Read(buf)
		// if err == io.EOF {
		// 	fmt.Println("EOF")
		// 	break
		// }
		// if err != nil {
		// 	fmt.Println("recv buf err:", err)
		// 	continue
		// }

		// 拆包
		dp := NewDataPack()

		// 读取客户端发送的 msg head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("Read msg headData err:", err)
			break
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("Unpack head err:", err)
			break
		}

		// 根据 dataLen 读取 data
		var data []byte
		if msg.GetLen() > 0 {
			data = make([]byte, msg.GetLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("Read msg data err:", err)
				break
			}
		}
		msg.SetData(data)

		// 得到 Request
		req := Request{
			conn: c,
			msg:  msg,
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

func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send message")
	}

	// 将 data 封包
	msg := NewMessage(msgID, data)

	dp := NewDataPack()
	binaryMsg, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack msg err:", err)
		return errors.New("pack msg err")
	}

	// 将数据发送给客户端
	_, err = c.Conn.Write(binaryMsg)
	if err != nil {
		fmt.Println("write binaryMsg err:", err)
		return errors.New("write binaryMsg err")
	}

	return nil
}
