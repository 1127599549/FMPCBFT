package main

type requester struct {
	//requesterID
	requesterID string
	//节点监听地址
	addr string
	//节点私钥
	priKey []byte
	//节点公钥
	pubKey []byte
	//proposers 对应的 ID 值
	proposers []string
}

type node struct {
	//节点ID
	nodeID string
	//节点监听地址
	addr string
	//节点私钥
	priKey []byte
	//节点公钥
	pubKey []byte
	//是否作为 proposer
	isProposer bool
	//对应的 primary verifiers
	priVerifiers []string
	//对应的 secondary verifiers
	secVerifiers []string
}

type Request struct {
	Message
	Timestamp int64
	//相当于 requesterID
	RequesterAddr string
}

type Message struct {
	Content string
	ID      int
}
