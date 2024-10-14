package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 1.连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("connect server error:" + err.Error())
		return
	}

	// 2.发数据
	for {
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println("client write error:" + err.Error())
			return
		}

		buf := make([]byte, 512)
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client read error:" + err.Error())
			return
		}

		fmt.Printf("callback %s, %d\n", buf, count)
		time.Sleep(time.Second * 2)
	}
}
