package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	go func() {
		dial, err := net.DialTimeout("tcp", "10.149.152.35:9099", 100*time.Second)
		if err != nil {
			panic(err)
		}
		println("connect ok")
		var buffer []byte
		for {
			if bs, e := dial.Read(buffer); e == nil || bs > 0 {
				fmt.Println(string(buffer))
			}
			dial.Write([]byte("Ping"))
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		fmt.Println("Hello World")
	}()
	fmt.Println("Hello World")
	select {}
}
