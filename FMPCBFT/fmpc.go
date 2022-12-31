package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type fmpc struct {
	//节点信息
	node node
	//每笔请求自增序号
	sequenceID int
	//锁
	lock sync.Mutex
	//临时消息池，消息摘要对应消息本体
	messagePool map[string]Request
}

func NewFMPC(nodeID, addr string) *fmpc {
	f := new(fmpc)
	f.node.nodeID = nodeID
	f.node.addr = addr
	f.node.priKey = getPriKey(nodeID)
	f.node.pubKey = getPubKey(nodeID)
	f.sequenceID = 0
	f.messagePool = make(map[string]Request)
	return f
}

func (f *fmpc) handleRequest(data []byte) {
	// 待实现
	fmt.Println(data)
}

func (f *fmpc) sequenceIDAdd() {
	f.lock.Lock()
	f.sequenceID++
	f.lock.Unlock()
}

//传入节点编号，获取对应的公钥
func getPubKey(nodeID string) []byte {
	key, err := ioutil.ReadFile("Keys/" + nodeID + "/" + nodeID + "_PubKey")
	if err != nil {
		log.Panic(err)
	}
	return key
}

//传入节点编号，获取对应的私钥
func getPriKey(nodeID string) []byte {
	key, err := ioutil.ReadFile("Keys/" + nodeID + "/" + nodeID + "_PriKey")
	if err != nil {
		log.Panic(err)
	}
	return key
}
