package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	UPort int    = 22010
	QHost string = "www.baidu.com"
)

func main() {
	fmt.Println("~~~ dns client ~~~")

	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	commandline()
	log.Printf("--cmd--line-- uport:%v", UPort)

	// UseUDP()
	// hoststr := "www.baidu.com"
	list := DNSQuery("223.5.5.5", QHost)
	log.Printf("host:%v -dns-> iplist:%v", QHost, list)
}

func commandline() {
	flag.IntVar(&UPort, "uport", UPort, "srv-port: 服务器UDP端口")
	flag.StringVar(&QHost, "query", QHost, "需要查询的domain-name")

	flag.Parse()
}
