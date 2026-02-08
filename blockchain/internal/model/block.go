package model

import (
	"blockchain/blockchain/internal/hash"
	"strings"
	"time"
)

const Difficulty = 4

type Block struct {
	Index        int    `json:"index"`
	Nonce        int    `json:"nonce"`
	Timestamp    string `json:"timestamp"`
	Data         string `json:"data"`
	PreviousHash string `json:"previous_hash"`
	Hash         string `json:"hash"`
}

func mine(index int, data, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().UTC().Format(time.RFC3339),
		Data:         data,
		PreviousHash: previousHash,
	}

	prefix := strings.Repeat("0", Difficulty)
	for {
		hashCode := hash.BlockHash(block.Index, block.Nonce, block.Timestamp, block.Data, block.PreviousHash)
		if strings.HasPrefix(hashCode, prefix) {
			block.Hash = hashCode
			return block
		}
		block.Nonce++
	}
}

func ValidateBlock(block *Block) bool {
	hashCode := hash.BlockHash(block.Index, block.Nonce, block.Timestamp, block.Data, block.PreviousHash)
	return hashCode == block.Hash && strings.HasPrefix(hashCode, strings.Repeat("0", Difficulty))
}

func MineGenesisBlock() *Block {
	return mine(0, "Genesis Block", "0")
}

func MineBlock(data string, previousBlock *Block) *Block {
	return mine(previousBlock.Index+1, data, previousBlock.Hash)
}
