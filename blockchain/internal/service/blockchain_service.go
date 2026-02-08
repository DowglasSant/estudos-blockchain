package service

import (
	"blockchain/blockchain/internal/model"
	"blockchain/blockchain/internal/repository"
)

type BlockchainService struct {
	repo *repository.BlockchainRepository
}

func NewBlockchainService(repo *repository.BlockchainRepository) *BlockchainService {
	svc := &BlockchainService{repo: repo}
	svc.repo.AddBlock(model.MineGenesisBlock())
	return svc
}

func (s *BlockchainService) MineBlock(data string) *model.Block {
	block := model.MineBlock(data, s.repo.LastBlock())
	s.repo.AddBlock(block)
	return block
}

func (s *BlockchainService) GetChain() []*model.Block {
	return s.repo.GetChain()
}

func (s *BlockchainService) IsValid() bool {
	return s.repo.IsValid()
}
