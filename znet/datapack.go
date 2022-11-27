// datapack 提供解决 TCP 粘包问题的解决方案，实现 Message 传输

package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"my-zinx/utils"
	"my-zinx/ziface"
)

func NewDataPack() *DataPack {
	return &DataPack{}
}

// DataPack 提供对 Message 封包和拆包的方法
type DataPack struct{}

func (d *DataPack) GetHeadLen() uint32 {
	return 8 // dataLen = 4, dataID = 4, 因此 head 长度为 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 写 DataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetLen()); err != nil {
		return nil, err
	}

	// 写 ID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetID()); err != nil {
		return nil, err
	}

	// 写 DATA
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (d *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建 IOReader
	reader := bytes.NewReader(binaryData)

	// 创建 Message
	msg := &Message{}

	// 读 DataLen
	if err := binary.Read(reader, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读 ID
	if err := binary.Read(reader, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	// 判断 DataLen 是否超出最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data received")
	}

	return msg, nil
}
