package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// requester 使用 tcp 监听
func requesterTcpListen() {
	listen, err := net.Listen("tcp", requesterAddr)
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		b, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(string(b))
	}
}

//节点使用 tcp 监听
func (f *fmpc) tcpListen() {
	listen, err := net.Listen("tcp", f.node.addr)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("节点开启监听，地址： %s \n", f.node.addr)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		b, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Panic(err)
		}
		f.handleRequest(b)
	}
}

//使用 tcp 发送消息
func tcpDial(context []byte, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("connect error", err)
		return
	}
	_, err = conn.Write(context)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
