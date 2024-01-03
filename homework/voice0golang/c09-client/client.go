package main

import (
	"c09comm"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	Port   int    = 22009
	SrcDir string = "../client-dir/"
	Fname  string = ""
	OpType int    = 1
)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("~~~ c09-client ~~~")

	commandline()
	log.Printf("--line-- srvport:%v datadir:%v optype:%v fname:%v ",
		Port,
		SrcDir,
		OpType,
		Fname)

	//提前创建文件夹
	SrcDir += "/"
	os.MkdirAll(SrcDir, 0644)

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%v", Port))
	if err != nil {
		log.Fatal(err)
	}

	var opcode uint32
	if OpType == c09comm.Op_Upload {
		opcode = c09comm.Op_Upload
	} else {
		opcode = c09comm.Op_Download
	}

	//写消息头
	header := make([]byte, 1024)
	binary.LittleEndian.PutUint32(header, opcode)
	bufname := []byte(Fname)
	binary.LittleEndian.PutUint32(header[4:], uint32(len(bufname)))
	_ = append(header[8:8], bufname...)

	n, err := conn.Write(header)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("write n:%v", n)

	if OpType == c09comm.Op_Upload {
		err := c09comm.SendFile(conn, SrcDir+Fname)
		log.Printf("upload err:%v file:%v", err, Fname)
	} else {
		err := c09comm.RecvFile(conn, SrcDir+Fname)
		log.Printf("download err:%v file:%v", err, Fname)
	}

	// time.Sleep(time.Second)
}

func commandline() {
	flag.IntVar(&Port, "port", Port, "dest-srv-port: 服务器端口")
	flag.StringVar(&SrcDir, "srcDir", SrcDir, "client data dir: 客户端数据文件夹")
	flag.IntVar(&OpType, "optype", OpType, "操作类型：1-上传，2-下载")
	flag.StringVar(&Fname, "fname", Fname, "上传或者下载的文件名")

	flag.Parse()
}
