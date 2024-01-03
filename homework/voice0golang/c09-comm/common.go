package c09comm

import (
	"net"
	"os"
)

// Head:1024byte [0-4]:optype [4-8]:len(fname) [8-xx]:fname [xx-1024]:0-padding [1024-xxx]:filedata
const (
	Op_Upload   = 1 //上传操作
	Op_Download = 2 //下载操作
	Header_Size = 1024
)

func SendFile(conn net.Conn, fname string) error {
	buf, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	ipos := 0
	for {
		n, err := conn.Write(buf[ipos:])
		ipos += n
		if err != nil {
			return err
		}
		if ipos == len(buf) {
			break
		}
	}

	return nil
}

func RecvFile(conn net.Conn, fname string) error {
	fobj, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fobj.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n != 0 {
			fobj.Write(buf[:n])
		}
		if err != nil { //对方断开连接，表示接受结束了
			break
		}
	}

	return err
}
