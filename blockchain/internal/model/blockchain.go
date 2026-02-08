package model

type Blockchain struct {
	Chain []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Chain: []*Block{},
	}
}

func (bc *Blockchain) AddBlock(block *Block) {
	bc.Chain = append(bc.Chain, block)
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) IsValid() bool {
	for i, block := range bc.Chain {
		if !ValidateBlock(block) {
			return false
		}

		if i > 0 && block.PreviousHash != bc.Chain[i-1].Hash {
			return false
		}
	}
	return true
}
