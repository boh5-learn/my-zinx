package utils

import (
	"encoding/json"
	"fmt"
	"my-zinx/ziface"
	"os"
)

// init 初始化 GlobalObject
func init() {
	// 默认配置
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0",
		Port:           8999,
		Name:           "ZinxServerApp",
		Version:        "v0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 尝试从 conf/zinx.json 加载配置
	GlobalObject.Reload()
}

// GlobalObj 存储所有 Zinx 的全局参数，供其他模块使用
// 一些参数通过 zinx.json 由用户配置
type GlobalObj struct {
	// server 配置
	TcpServer ziface.IServer // 当前 Zinx 全局 Server 对象
	Host      string         // 监听 IP
	Port      int            // 监听端口
	Name      string         // 服务器名称

	// Zinx 配置
	Version        string // Zinx 版本号
	MaxConn        int    // 最大连接数
	MaxPackageSize uint32 // 数据包一次最大读多少字节
}

// GlobalObject 定义一个全局的对外的 GlobalObj
var GlobalObject *GlobalObj

// Reload 去 zinx.json 加载配置
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(fmt.Sprintln("Read zinx.json file err:", err))
	}

	// 将 json 数据解析到 GlobalObj 中
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(fmt.Sprintln("Load zinx.json data err:", err))
	}
}
