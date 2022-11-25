package znet

import (
	"my-zinx/ziface"
)

// BaseRouter router 基类，实现 router 时，先嵌入 BaseRouter，然后再根据需要重写方法即可
type BaseRouter struct{}

// 这里之所以 BaseRouter 的方法都为空，
// 是因为有的 Router 不需要所有方法，
// 这样，只需重写需要的方法就可以了。
// 否则，如果用接口的话，必须实现三个方法

func (b *BaseRouter) PreHandle(request ziface.IRequest) {}

func (b *BaseRouter) Handle(request ziface.IRequest) {}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {}
