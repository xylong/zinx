package ziface

// IRequest 将客户端请求的连接信息和请求数据包装到一个request中
type IRequest interface {
	// GetConnection 获取连接
	GetConnection() IConnection
	// GetData 获取请求数据
	GetData() []byte
}
