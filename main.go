package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
)

func main () {
	fmt.Println("start listening")

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		t, _ := conn.(*net.TCPConn)
		f, _ := t.File()
		log.Printf("new connection fd:%d\n",f.Fd())
		if err != nil {
			log.Fatal(err)
		}
		go handleRpc(conn)
	}
}

func handleRpc(conn net.Conn) {
	header := make([]byte, 4)
	n, err := conn.Read(header)
	if err != nil {
		log.Fatal(err)
	}
	// TLV L=int32=4B
	if n < 4 {
		log.Fatal(fmt.Sprintf("n:%d < 4",n))
	}

	size := int32(binary.BigEndian.Uint32(header))
	log.Printf("body size = %d\n",size)

	buffer := make([]byte, size)
	n, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	if n < int(size) {
		log.Fatal(fmt.Sprintf("n:%d < size:%d",n,size))
	}

	msg := string(buffer)
	log.Printf("body = %s\n",msg)
	msgArr := strings.Split(msg, " ")
	if len(msgArr) != 3 {
		log.Fatal(fmt.Sprintf("msg arr:%v length is wrong",msgArr))
	}

	var ans int
	a, err := strconv.ParseInt(msgArr[1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.ParseInt(msgArr[2], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	switch msgArr[0] {
	case "ADD":
		ans = int(a+b)
	case "SUB":
		ans = int(a-b)
	case "MUL":
		ans = int(a*b)
	case "DIV":
		ans = int(a/b)
	default:
		log.Fatal(fmt.Sprintf("not support :%s",msgArr[0]))
	}

	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(ans))
	_, err = conn.Write(res)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
