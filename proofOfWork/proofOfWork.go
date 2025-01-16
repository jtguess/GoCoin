package proofofwork

import (
	"goCoin/blockchain"
	"math/big"
)

const targetBits = 24

type ProofOfWork struct {
	block  *blockchain.Block
	target *big.Int
}

func NewProofOfWork(b *blockchain.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}
