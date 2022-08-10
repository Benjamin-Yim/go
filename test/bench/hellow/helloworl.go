package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func main() {
	go func() {
		listen, err := net.Listen("tcp", "0.0.0.0:20000")
		if err != nil {
			fmt.Println("listen failed, err:", err)
			return
		}
		for {
			//2.接收客户端的链接
			conn, err := listen.Accept()
			if err != nil {
				fmt.Printf("accept failed, err:%v\n", err)
				continue
			}
			//3.开启一个Goroutine，处理链接
			go process(conn)
		}
	}()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Println("=============================================================goroutine nums:", runtime.NumGoroutine())
		}
	}()
	fmt.Println("Hello World")
	select {}
}

// 处理请求，类型就是net.Conn
func process(conn net.Conn) {

	//处理结束后关闭链接
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v", err)
			break
		}
		fmt.Printf("recv from client, content:%v\n", string(buf[:n]))
	}

}
