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
	// 当前连接绑定的处理业务的方法
	handle ziface.HandleFunc
	// 退出通知
	ExitChan chan struct{}
}

func NewConnection(conn *net.TCPConn, connID uint64, handle ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		handle:   handle,
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
		count, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("reader receive buf error:" + err.Error())
			continue
		}

		if err = c.handle(c.Conn, buf, count); err != nil {
			fmt.Printf("connID=%d handle error:", err.Error())
			break
		}
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

func (c *Connection) Send() error {
	return nil
}
