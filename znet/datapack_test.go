package znet

import (
	"io"
	"net"
	"testing"
)

// TestDataPack 测试 DataPack 的拆包和封包
func TestDataPack(t *testing.T) {
	addr := "127.0.0.1:7777"
	// 模拟服务器
	// 1. 创建 Socket
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatal("Server listen err:", err)
	}

	// 2. 创建 goroutine 复制处理客户端业务
	go func() {
		// 从客户端读取数据，拆包
		conn, err := listener.Accept()
		if err != nil {
			t.Error("Server Accept err:", err)
			return
		}

		// 处理客户端请求
		go func(conn net.Conn) {
			// 拆包
			dp := NewDataPack()
			for {
				// 读头部
				head := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, head)
				if err != nil {
					t.Error("Read head failed, err:", err)
					return
				}

				// Unpack
				msgHead, err := dp.Unpack(head)
				if err != nil {
					t.Error("Unpack head err:", err)
					return
				}

				// 有 msg，则读 msg
				if msgHead.GetLen() > 0 {
					msg := msgHead.(*Message)
					// 开辟 msg 空间
					msg.Data = make([]byte, msg.GetLen())

					// 读 msg
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						t.Error("Read msg err:", err)
						return
					}

					// 打印完整消息
					t.Logf("Recv ID: %d, DataLen: %d, Data: %s", msg.GetID(), msg.GetLen(), msg.GetData())
				}
			}
		}(conn)
	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal("Client dial err:", err)
	}

	dp := NewDataPack()

	// 封装 2 个 Message 一起发送
	msg1 := &Message{
		ID:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		t.Fatal("Client pack msg1 err:", err)
	}

	msg2 := &Message{
		ID:      2,
		DataLen: 11,
		Data:    []byte{'H', 'e', 'l', 'l', 'o', ',', 'z', 'i', 'n', 'x', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		t.Fatal("Client pack msg2 err:", err)
	}

	// 两个包粘在一起
	sendData1 = append(sendData1, sendData2...)

	// 一起发送
	_, err = conn.Write(sendData1)
	if err != nil {
		t.Error("Client send msg err:", err)
	}

	select {}
}
