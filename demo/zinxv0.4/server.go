package main

import (
	"fmt"
	"my-zinx/ziface"
	"my-zinx/znet"
)

/*
基于 Zinx 框架来开发的服务器应用程序
*/

func main() {
	// 1. 创建 server
	s := znet.NewServer("[zinx v0.2]")
	// 2. 绑定 Router
	s.AddRouter(&PingRouter{})
	// 3. 启动 server
	s.Serve()
}

// PingRouter ping 测试自定义 Router
type PingRouter struct {
	znet.BaseRouter
}

// PreHandle Test PreHandle
func (b *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	if err != nil {
		fmt.Println("PreHandle err:", err)
		return
	}
}

// Handle Test Handle
func (b *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("Handle err:", err)
		return
	}
}

// PostHandle Test PostHandle
func (b *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil {
		fmt.Println("PostHandle err:", err)
		return
	}
}
