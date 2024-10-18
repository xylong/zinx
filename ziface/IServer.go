package ziface

type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Run 运行服务器
	Run()
	// AddRouter 添加路由，供客户端连接处理使用
	AddRouter(IRouter)
}
