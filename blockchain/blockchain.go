package blockchain

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

// singleton 패턴
var b *blockchain

func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}
