package main

import (
	"fmt"

	"github.com/hmokingbird/LULChain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.AddBlock("Test 1")
	bc.AddBlock("Test 2")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.Prev)
		fmt.Println()
	}
}
