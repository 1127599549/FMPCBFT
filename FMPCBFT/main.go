package main

import (
	"log"
	"os"
)

//requester的监听地址
var requesterAddr = "127.0.0.1:8888"

// const requesterCount = 1

const nodeCount = 4

//节点池，主要用于存储监听的地址
var nodeTable map[string]string

func main() {
	//为四个节点生成公私钥
	generateKeys()
	//requesterTable = map[string]string{
	//	"R1": "127.0.0.1:8888",
	//}
	nodeTable = map[string]string{
		"N0": "127.0.0.1:8000",
		"N1": "127.0.0.1:8001",
		"N2": "127.0.0.1:8002",
		"N3": "127.0.0.1:8003",
	}
	if len(os.Args) != 2 {
		log.Panic("输入的参数有误！ ")
	}
	nodeID := os.Args[1]
	if nodeID == "requester" {
		requesterSendMessageAndListen() //启动客户端程序
	} else if addr, ok := nodeTable[nodeID]; ok {
		f := NewFMPC(nodeID, addr)
		go f.tcpListen()
	} else {
		log.Fatal("无此节点编号！")
	}
	select {}
}
