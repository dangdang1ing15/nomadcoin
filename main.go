package main

import (
	"crypto/sha256"
	"fmt"
)

// 데이터를 오직 block에다가만 저장하는 blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

// block을 hash해보기

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash)) //string 데이터를 byte 형식인 slice로 바꿔줌
	hexHash := fmt.Sprintf("%x", hash)                                       // 바이트를 16진수로 출력하고 반환시킴
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}

// 스트링을 바꿔준다는 거에서 번거롭, 자동화가 안되어있어 블록마다 복붙해야하는 문제 있음, 해결예정
