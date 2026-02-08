package repository

import "blockchain/blockchain/internal/model"

type BlockchainRepository struct {
	blockchain *model.Blockchain
}

func NewBlockchainRepository() *BlockchainRepository {
	return &BlockchainRepository{
		blockchain: model.NewBlockchain(),
	}
}

func (r *BlockchainRepository) AddBlock(block *model.Block) {
	r.blockchain.AddBlock(block)
}

func (r *BlockchainRepository) LastBlock() *model.Block {
	return r.blockchain.LastBlock()
}

func (r *BlockchainRepository) GetChain() []*model.Block {
	return r.blockchain.Chain
}

func (r *BlockchainRepository) IsValid() bool {
	return r.blockchain.IsValid()
}
