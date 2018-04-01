package core

const MAX_QUEUE_SIZE int = 100

type BlockChain struct {
	Chain [MAX_QUEUE_SIZE]*Block
	Front int
	Rear  int
}

func NewBlockChain() *BlockChain {
	blockchain := BlockChain{
		Front: 0,
		Rear:  0,
	}

	return &blockchain
}

func (blockchain *BlockChain) Push(block *Block) {

	front := &blockchain.Front
	rear := &blockchain.Rear

	blockchain.Chain[*rear] = block

	*rear = (*rear + 1) % MAX_QUEUE_SIZE
	if *rear == *front {
		*front = (*front + 1) % MAX_QUEUE_SIZE
	}

}

func (blockchain *BlockChain) GetLastBlock() *Block {
	return blockchain.Chain[blockchain.Rear-1]
}

func (blockchain *BlockChain) CreateChain() *Block {
	block := NewGenesisBlock()
	blockchain.Push(block)

	return block
}

func (blockchain *BlockChain) NewChain(transactions []Transaction) *Block {

	prevblock := blockchain.GetLastBlock()
	block := NewBlock(prevblock.Header, transactions)
	blockchain.Push(block)

	return block

}
