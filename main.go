package main

import (
	"fmt"
	"goCoin/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Sent one GoCoin to Connor")
	bc.AddBlock("Sent 2 more GoCoin to Connor")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
