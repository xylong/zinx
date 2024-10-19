package utils

import (
	"encoding/json"
	"os"
	"zinx/ziface"
)

// GlobalObj 全局参数
type GlobalObj struct {
	TcpServer ziface.IServer
	Host      string
	TcpPort   int    // tcp监听端口
	Name      string // 服务名称

	Version        string // 版本号
	MaxConn        int    // 服务允许的最大连接数
	MaxPackageSize uint64 // 支持的数据包最大值
}

// Reload 加载用户自定义配置
func (o *GlobalObj) Reload() {
	bt, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(bt, GlobalObject); err != nil {
		panic(err)
	}
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Name:           "Zinx",
		Version:        "v0.1",
		TcpPort:        8899,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
}
