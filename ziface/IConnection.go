package ziface

import "net"

// IConnection 连接模块
type IConnection interface {
	// Start 启动链接Start()
	Start()
	// Stop 停止链接
	Stop()
	// GetTCPConnection 获取当前链接的conn对象(套接字)
	GetTCPConnection() *net.TCPConn
	// GetConnID 得到链接🆔
	GetConnID() uint64
	// GetRemoteAddr 得到客户端链接的地址和端口
	GetRemoteAddr() net.Addr
	// Send 发送数据
	Send([]byte) error
}

// HandleFunc 处理连接业务方法
type HandleFunc func(*net.TCPConn, []byte, int) error
