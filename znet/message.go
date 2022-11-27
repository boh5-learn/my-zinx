package znet

type Message struct {
	ID      uint32 // 消息 ID
	DataLen uint32 // 消息长度
	Data    []byte // 消息内容
}

func (m *Message) GetID() uint32 {
	return m.ID
}

func (m *Message) GetLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetID(id uint32) {
	m.ID = id
}

func (m *Message) SetLen(length uint32) {
	m.DataLen = length
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
