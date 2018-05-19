package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prev := bc.Blocks[len(bc.Blocks)-1]
	new := NewBlock(data, prev.Hash)
	bc.Blocks = append(bc.Blocks, new)
}

func FirstBlock() *Block {
	return NewBlock("Genesis", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{FirstBlock()}}
}
