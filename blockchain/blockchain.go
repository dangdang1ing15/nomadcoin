package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

// #1 데이터를 오직 block에다가만 저장하는 blockchain

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHast,omitempty"`
	Height   int    `josn:"height"`
}

// #2 blockchain struct 생성해 block들의 slice를 가지고 있다 정의

type blockchain struct {
	blocks []*Block
}

// singleton 패턴
var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {

	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockchain().blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

var ErrNotFound = errors.New("block not fount")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil
}
