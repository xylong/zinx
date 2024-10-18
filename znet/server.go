package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	// Name 服务器名称
	Name string
	// IPVersion ip版本
	IPVersion string
	// IP 服务器监听的ip
	IP string
	// Port 服务器监听端口
	Port int
	// Router 路由
	Router ziface.IRouter
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8899,
		Router:    nil,
	}
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("start tcp on %s:%d\n", s.IP, s.Port)

	go func() {
		// 1.获取tcp连接地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:" + err.Error())
			return
		}

		// 2.监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp addr error:" + err.Error())
			return
		}

		// 3.阻塞等待客户端连接，处理客户端业务(读写)
		var cid uint64
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept tcp error" + err.Error())
				continue
			}

			// 和connection模块绑定
			dealConn := NewConnection(conn, cid, s.Router)
			go dealConn.Start()
			cid++
		}
	}()

	// todo 扩展
}

func (s *Server) Stop() {

}

func (s *Server) Run() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}
