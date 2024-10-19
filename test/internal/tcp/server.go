package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (r *PingRouter) Before(request ziface.IRequest) {
	fmt.Println("------before-----")
	request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
}

func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("====handle====")
	request.GetConnection().GetTCPConnection().Write([]byte("ping..."))
}

func (r *PingRouter) After(request ziface.IRequest) {
	fmt.Println("~~~~after~~~~~")
	request.GetConnection().GetTCPConnection().Write([]byte("after ping..."))
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Run()
}
