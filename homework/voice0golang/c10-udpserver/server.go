package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	UPort int = 22010
)

func main() {
	fmt.Println("~~~ c10: udp server ~~~")

	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	commandline()
	log.Printf("--cmd--line-- uport:%v", UPort)

	udpaddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", UPort))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1600) // max-onepack-len:1472
	for {
		n, cli, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("recv-cli:%v data:%v", cli, string(buf[:n]))

		conn.WriteToUDP([]byte("success"), cli)
	}
}

func commandline() {
	flag.IntVar(&UPort, "uport", UPort, "srv-port: 服务器UDP端口")

	flag.Parse()
}
