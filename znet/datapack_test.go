package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {

	// 模拟服务端
	listener, err := net.Listen("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("server listen error:", err.Error())
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept error:", err.Error())
			}

			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					// 第一次读，读取head
					header := make([]byte, dp.GetHeadLength())
					_, err := io.ReadFull(conn, header)
					if err != nil {
						fmt.Println("read head error:", err.Error())
						break
					}
					msg, err := dp.Unpack(header)
					if err != nil {
						fmt.Println("unpack error:", err.Error())
						return
					}

					// 第二次读，根据head信息读取data
					if msg.GetMsgLen() > 0 {
						data := make([]byte, msg.GetMsgLen())
						_, err = io.ReadFull(conn, data)
						if err != nil {
							fmt.Println("second unpack error:", err.Error())
							return
						}
						msg.SetData(data)
					}

					fmt.Println("-->msgId=", msg.GetMsgId(), "len=", msg.GetMsgLen(), "data=", string(msg.GetData()))
				}

			}(conn)
		}
	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("connect server error:" + err.Error())
		return
	}
	// 模拟粘包，封装两个包一起发送
	dp := NewDataPack()
	msg1 := &Message{
		Id:     1,
		Length: 4,
		Data:   []byte{'z', 'i', 'n', 'x'},
	}
	pack1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1 error:", err.Error())
		return
	}

	msg2 := &Message{
		Id:     2,
		Length: 11,
		Data:   []byte("hello world"),
	}
	pack2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack msg2 error:", err.Error())
		return
	}

	// 将两个包粘在一起
	pack1 = append(pack1, pack2...)
	conn.Write(pack1)

	select {}
}
