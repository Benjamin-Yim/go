package helloworld

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func BenchmarkHelloWorld(b *testing.B) {
	b.ReportAllocs()
	dial, err := net.DialTimeout("tcp", "192.168.123.128:9090", 100*time.Second)
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

}
