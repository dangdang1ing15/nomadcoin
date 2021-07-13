package main

import (
	"crypto/sha256"
	"fmt"
)

// #1 데이터를 오직 block에다가만 저장하는 blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

// #2 blockchain struct 생성해 block들의 slice를 가지고 있다 정의

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLashHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLashHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("PrevHash: %s\n", block.prevHash)
	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}

// #1 스트링을 바꿔준다는 거에서 번거롭, 자동화가 안되어있어 블록마다 복붙해야하는 문제 있음, 해결예정
// #2 main 함수에서 분리하는 작업 진행, block 배열을 만들고, 이전 해시를 리턴하는 함수, 블록을 만드는 함수, 블록의 정보를 출력하는 함수, 그리고 메인에선 블록을 추가하는 명령어들로 구현되어있음. 그러나 더하는 함수쪽에 총 3개지 기능 (블록 만들기, 해싱, 해싱값 slice에 추가)있어 분리예정
