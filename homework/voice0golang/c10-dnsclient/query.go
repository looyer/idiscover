package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	DNSPort = 53 //dns服务器默认端口

	Type_RR_IP    = 1
	Type_RR_CName = 5
)

// doc: https://en.wikipedia.org/wiki/Domain_Name_System  [DNS message format]
type DNSHeader struct {
	ID            uint16
	Flags         uint16 //Q(0)R(1):1bit | OPCODE:4bit | AA:1bit | TC:1bit | RD:1bit | RA:1bit | Z(reserved=0):3bit | RCODE:4bit
	NumQuestion   uint16 //查询问题的数量
	NumAnswer     uint16 //返回答案的数量
	NumAuthority  uint16
	NumAdditional uint16
}

func (p *DNSHeader) SetFlags(qr, op, aa, tc, rd, ra, rcode uint16) {
	p.Flags = qr<<15 | op<<11 | aa<<10 | tc<<9 | rd<<8 | ra<<7 | rcode
}

func DNSQuery(dnsServe, host string) (ans []string) {
	conn, err := net.Dial("udp", fmt.Sprintf("%v:%v", dnsServe, DNSPort))
	if err != nil {
		log.Fatal(err)
	}

	var (
		buf bytes.Buffer
	)

	//写入DNS请求头
	header := DNSHeader{ID: 0xFFFF, NumQuestion: 1}
	header.SetFlags(0, 0, 0, 0, 0, 0, 0)
	binary.Write(&buf, binary.BigEndian, header)

	//写入问题
	qs := questionSection(host)
	buf.Write(qs)

	log.Print(buf)

	//请求和接收数据
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	recvbuf := make([]byte, 1500)
	n, err := conn.Read(recvbuf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(n, recvbuf[:n])

	return parseReply(recvbuf[:n])
}

func questionSection(host string) []byte {
	buf := make([]byte, 0, 256)

	vSegs := strings.Split(host, ".")
	for _, str := range vSegs {
		seg := []byte(str)
		buf = append(buf, byte(len(seg)))
		buf = append(buf, seg...)
	}
	//末尾0
	buf = append(buf, 0)
	//TypeOfRR //主机地址:0x01 IPV6地址:0x1c
	buf = binary.BigEndian.AppendUint16(buf, 0x01)
	//ClassCode
	buf = binary.BigEndian.AppendUint16(buf, 0x01)

	return buf
}

func parseReply(buf []byte) (ans []string) {
	ans = make([]string, 0)

	if 12 > len(buf) { //不足头部
		return
	}
	header := DNSHeader{}
	header.ID = binary.BigEndian.Uint16(buf)
	header.Flags = binary.BigEndian.Uint16(buf[2:])
	header.NumQuestion = binary.BigEndian.Uint16(buf[4:])
	header.NumAnswer = binary.BigEndian.Uint16(buf[6:])
	header.NumAuthority = binary.BigEndian.Uint16(buf[8:])
	header.NumAdditional = binary.BigEndian.Uint16(buf[10:])

	log.Print(header)

	//去掉问题部分
	body := buf[12:]
	ipos := uint16(0) //移动游标
	for i := 0; i < int(header.NumQuestion); i++ {
		for {
			if ipos+1 > uint16(len(body)) { //不足以解析
				return
			}
			n := uint16(body[ipos])
			if n == 0 {
				ipos += 5 //end:0(1bit) + typeofRR(2bit) + class(2bit)
				break
			}
			ipos += 1 + n
		}
	}

	//解析答案部分
	for i := 0; i < int(header.NumAnswer); i++ {
		//域名:2byte 类型:2byte 类:2byte 生存时间:4byte 资源数据长度:2byte 资源数据...
		if ipos+(2+2+2+4+2) > uint16(len(body)) { //不足以解析一条答案的头
			return
		}
		reslen := binary.BigEndian.Uint16(body[ipos+2+2+2+4:])

		if ipos+(2+2+2+4+2)+reslen > uint16(len(body)) {
			return
		}

		str := parseAnswer(body[ipos : ipos+(2+2+2+4+2)+reslen])
		ans = append(ans, str)

		ipos += (2 + 2 + 2 + 4 + 2) + reslen
	}

	return
}

// 解析一条answer数据
func parseAnswer(rr []byte) string {
	ctype := binary.BigEndian.Uint16(rr[2:])
	reslen := binary.BigEndian.Uint16(rr[2+2+2+4:])
	ipos := 2 + 2 + 2 + 4 + 2

	log.Print(rr)

	if ctype == Type_RR_IP {
		if reslen == 4 {
			return fmt.Sprintf("%v.%v.%v.%v", rr[ipos], rr[ipos+1], rr[ipos+2], rr[ipos+3])
		}

	} else if ctype == Type_RR_CName {
		cname := ""
		for {
			if ipos > len(rr) {
				return strings.TrimSuffix(cname, ".")
			}
			n := int(rr[ipos])
			ipos++
			if ipos+n > len(rr) {
				return strings.TrimSuffix(cname, ".")
			}
			cname += string(rr[ipos:ipos+n]) + "."
			ipos += n
		}
	}
	return ""
}
