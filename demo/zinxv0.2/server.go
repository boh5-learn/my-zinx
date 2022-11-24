package main

import "my-zinx/znet"

/*
基于 Zinx 框架来开发的服务器应用程序
*/

func main() {
	// 1. 创建 server
	s := znet.NewServer("[zinx v0.2]")
	// 2. 启动 server
	s.Serve()
}
