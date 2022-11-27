package ziface

// IDataPack 实现对 IMessage 的 TCP 封包，拆包方法，解决粘包问题
type IDataPack interface {
	// GetHeadLen 获取包头的长度
	GetHeadLen() uint32
	// Pack 封包方法
	Pack(msg IMessage) ([]byte, error)
	// Unpack 拆包方法
	Unpack(binaryData []byte) (IMessage, error)
}
