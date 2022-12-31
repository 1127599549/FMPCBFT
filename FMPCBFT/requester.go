package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

func requesterSendMessageAndListen() {
	//开启requester的本地监听（主要用来接收节点的reply信息）
	go requesterTcpListen()
	fmt.Printf("requester开启监听，地址: %s\n", requesterAddr)

	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("|  已进入FMPC测试Demo客户端，请启动全部节点后再发送消息！  ：）  |")
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("请在下方输入要存入节点的信息： ")
	//首先通过命令行获取用户输入
	stdReader := bufio.NewReader(os.Stdin)
	for {
		data, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}
		r := new(Request)
		r.Timestamp = time.Now().UnixNano()
		r.RequesterAddr = requesterAddr
		r.Message.ID = getRandom()
		//消息内容就是用户的输入
		r.Message.Content = strings.TrimSpace(data)
		br, err := json.Marshal(r)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(string(br))
		tcpDial(br, nodeTable["N0"])
	}
}

//返回一个十位数随机数，作为msgid
func getRandom() int {
	x := big.NewInt(10000000000)
	for {
		result, err := rand.Int(rand.Reader, x)
		if err != nil {
			log.Panic(err)
		}
		if result.Int64() > 1000000000 {
			return int(result.Int64())
		}
	}
}
