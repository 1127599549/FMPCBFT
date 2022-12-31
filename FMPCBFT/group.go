package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"github.com/yoseplee/vrf"
	_ "github.com/yoseplee/vrf"
	"github.com/yoseplee/vrf/sortition"
	_ "github.com/yoseplee/vrf/sortition"
	"log"
	"os"
	"sort"
	"strconv"
)

//如果当前目录下不存在目录Keys，则创建目录，并为各个节点生成椭圆曲线公私钥

func generateKeys() {
	if !isExist("./Keys") {
		fmt.Println("检测到还未生成公私钥目录，正在生成公私钥 ...")
		err := os.Mkdir("Keys", 0644)
		if err != nil {
			log.Panic()
		}

		for i := 0; i <= nodeCount; i++ {
			if !isExist("./Keys/N" + strconv.Itoa(i)) {
				err := os.Mkdir("./Keys/N"+strconv.Itoa(i), 0644)
				if err != nil {
					log.Panic()
				}
			}
			priKey, pubKey := getKeyPair()
			priFileName := "Keys/N" + strconv.Itoa(i) + "/N" + strconv.Itoa(i) + "_PriKey"
			file, err := os.OpenFile(priFileName, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			file.Write(priKey)

			pubFileName := "keys/N" + strconv.Itoa(i) + "/N" + strconv.Itoa(i) + "_PubKey"
			file2, err := os.OpenFile(pubFileName, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file2.Close()
			file2.Write(pubKey)
		}
		fmt.Println("已为节点们生成椭圆曲线公私钥")
	}
}

//判断文件或文件夹是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

//生成椭圆曲线公私钥
func getKeyPair() (priKey, pubKey []byte) {
	//生成公私钥
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	priKey = pri
	pubKey = pub
	return
}

//数字签名
//vrf.Prove

//签名验证
//vrf.Verify

//通过 vrf 选出 proposers
func (r requester) calProposers() {
	var ratio []float64
	for i := 0; i < nodeCount; i++ {
		pubKey := getPubKey("N" + strconv.Itoa(i))
		priKey := getPriKey("N" + strconv.Itoa(i))
		var m []byte = []byte("N" + strconv.Itoa(i))
		_, hash, err := vrf.Prove(pubKey, priKey, m)
		if err != nil {
			panic(err)
		}
		ration1 := sortition.HashRatio(hash)
		if sortition.Sortition(ration1) {
			ratio = append(ratio, ration1)
		}
	}
	if len(ratio) >= 3 {
		sort.Float64s(ratio)
	}
}

//通过 vrf 选出 proposer 对应的 主次 verifiers
