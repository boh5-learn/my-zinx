package ziface

// IServer 定义服务器接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()
	// AddRouter 路由功能：给当前服务提供注册 Router 的方法
	AddRouter(router IRouter)
}
