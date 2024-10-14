package znet

import (
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
	// 当前连接绑定的处理业务的方法
	handleApi ziface.HandleFunc
	// 退出通知
	exitChan chan struct{}
}

func NewConnection(conn *net.TCPConn, connID uint64, handle ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleApi: handle,
		isClosed:  false,
		exitChan:  make(chan struct{}, 1),
	}
}

func (c *Connection) Start() {

}

func (c *Connection) Stop() {

}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return nil
}

func (c *Connection) GetConnID() uint64 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return nil
}

func (c *Connection) Send() error {
	return nil
}
