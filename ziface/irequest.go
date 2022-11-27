package ziface

// IRequest 把客户端请求的 conn 和 data 包装到一个 Request 中
type IRequest interface {
	// GetConnection 得到当前连接
	GetConnection() IConnection
	// GetData 得到请求的数据
	GetData() []byte
	// GetID 得到请求的 ID
	GetID() uint32
}
