package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

// Connection 连接模块
type Connection struct {
	// Conn 当前链接套接字
	Conn *net.TCPConn
	// ConnID 连接🆔
	ConnID uint64
	// 是否关闭
	isClosed bool
	// Router 路由
	Router ziface.IRouter
	// 退出通知
	ExitChan chan struct{}
}

func NewConnection(conn *net.TCPConn, connID uint64, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan struct{}, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("reader running")
	defer fmt.Printf("connID=%d reader exit,remote address is %s\n", c.ConnID, c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("reader receive buf error:" + err.Error())
			continue
		}

		req := Request{
			conn: c,
			data: buf,
		}

		go func(request ziface.IRequest) {
			c.Router.Before(request)
			c.Router.Handle(request)
			c.Router.After(request)
		}(&req)
	}
}

// Start 启动链接
func (c *Connection) Start() {
	fmt.Printf("connection %d start\n", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Printf("connection %d stop\n", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	// 关闭socket连接
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint64 {
	return c.ConnID
}

// GetRemoteAddr 获取客户端ip、port
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
