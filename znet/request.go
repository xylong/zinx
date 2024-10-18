package znet

import "zinx/ziface"

type Request struct {
	// 客户端连接
	conn ziface.IConnection
	// 客户端请求数据
	data []byte
}

// GetConnection 获取客户端连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 获取请求数据
func (r *Request) GetData() []byte {
	return r.data
}
