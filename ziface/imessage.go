package ziface

// IMessage 封装请求的消息
type IMessage interface {
	// GetID 获取消息 ID
	GetID() uint32
	// GetLen 获取消息长度
	GetLen() uint32
	// GetData 获取消息内容
	GetData() []byte

	// SetID 设置消息 ID
	SetID(id uint32)
	// SetLen 设置消息长度
	SetLen(length uint32)
	// SetData 设置消息内容
	SetData(data []byte)
}
