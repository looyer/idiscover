package main

import (
	"c09comm"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"runtime/debug"
)

var (
	Port   int    = 22009
	SrcDir string = "../server-dir/"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("~~~ c09-server ~~~")

	commandline()
	log.Printf("--cmd--line-- port:%v srcdir:%v", Port, SrcDir)

	SrcDir += "/"
	os.MkdirAll(SrcDir, 0644)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", Port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handler(conn)
	}
}

func commandline() {
	flag.IntVar(&Port, "port", Port, "服务器端口")
	flag.StringVar(&SrcDir, "srcDir", SrcDir, "server data dir: 服务器数据文件夹")

	flag.Parse()
}

func handler(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("exception! panic:%v stack:%v", err, string(debug.Stack()))
		}
	}()
	defer conn.Close()

	ipos := 0 //
	header := make([]byte, c09comm.Header_Size)
	for {
		n, err := conn.Read(header[ipos:])
		ipos += n

		if err != nil && err != io.EOF {
			log.Panic(err)
		}
		if err == io.EOF || ipos == c09comm.Header_Size {
			break
		}
	}
	if ipos < c09comm.Header_Size {
		log.Panic(fmt.Errorf("conn:%v head-short-err! headsize:%v", conn.RemoteAddr().String(), ipos))
	}

	optype := binary.LittleEndian.Uint32(header)
	fnsize := binary.LittleEndian.Uint32(header[4:])
	fname := string(header[8 : 8+fnsize])

	if optype == c09comm.Op_Upload {
		err := c09comm.RecvFile(conn, SrcDir+fname)
		log.Printf("upload err:%v remote:%v fname:%v", err, conn.RemoteAddr(), fname)
	} else {
		err := c09comm.SendFile(conn, SrcDir+fname)
		log.Printf("download err:%v remote:%v fname:%v", err, conn.RemoteAddr(), fname)
	}
}
