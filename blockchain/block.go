package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	Prev      []byte
	Hash      []byte
}

func (b *Block) HashBlock() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	header := bytes.Join([][]byte{b.Prev, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(header)

	b.Hash = hash[:]
}

//Generates the new block based on the data
func NewBlock(data string, PreviousBlock []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Data: []byte(data), Prev: PreviousBlock, Hash: []byte{}}
	block.HashBlock()
	return block
}
