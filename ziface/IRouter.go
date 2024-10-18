package ziface

// IRouter 路由抽象
// 路由里的数据都是IRequest
type IRouter interface {
	// Before 处理业务之前钩子方法
	Before(IRequest)
	// Handle 处理业务的钩子
	Handle(IRequest)
	// After 处理业务之后钩子方法
	After(IRequest)
}
