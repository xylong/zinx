package znet

import "zinx/ziface"

// BaseRouter 基础路由
type BaseRouter struct {
}

func (r *BaseRouter) Before(request ziface.IRequest) {}
func (r *BaseRouter) Handle(request ziface.IRequest) {}
func (r *BaseRouter) After(request ziface.IRequest)  {}
