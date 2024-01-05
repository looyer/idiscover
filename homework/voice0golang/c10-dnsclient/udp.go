package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func UseUDP() {
	conn, err := net.Dial("udp", fmt.Sprintf("localhost:%v", UPort))
	if err != nil {
		log.Fatal(err)
	}

	//从控制台读取输入，发送
	go func() {
		line := make([]byte, 256)

		for {
			n, _ := os.Stdin.Read(line)
			conn.Write(line[:n])
		}
	}()

	buf := make([]byte, 1500)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("recv:%v ", string(buf[:n]))
	}
}
