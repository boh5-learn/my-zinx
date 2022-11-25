package ziface

// IRouter 路由, 传入 IRequest 进行处理
type IRouter interface {
	// PreHandle 处理 conn 业务之前的 hook
	PreHandle(request IRequest)
	// Handle 处理 conn 业务 hook
	Handle(request IRequest)
	// PostHandle 处理 conn 业务之后的 hook
	PostHandle(request IRequest)
}
