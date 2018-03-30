package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"log"
	"math"
	"time"
)

type BlockHeader struct {
	index         int64
	PrevBlockHash []byte
	MerkleRoot    []byte
	Timestamp     string
	Difficulty    int
	MiningReward  int
	Nonce         uint32
	Hash          []byte
}

type Block struct {
	Header       BlockHeader
	Transactions []Transaction
}

func (block *Block) Serialize() []byte {

	var exceptHash Block
	exceptHash = Block{
		Header: BlockHeader{
			index:         block.Header.index,
			PrevBlockHash: block.Header.PrevBlockHash,
			MerkleRoot:    block.Header.MerkleRoot,
			Timestamp:     block.Header.Timestamp,
			Difficulty:    block.Header.Difficulty,
			MiningReward:  block.Header.MiningReward,
			Nonce:         block.Header.Nonce,
			Hash:          []byte{0},
		},
		Transactions: block.Transactions,
	}

	var serialized bytes.Buffer
	encoder := gob.NewEncoder(&serialized)
	err := encoder.Encode(exceptHash)
	if err != nil {
		log.Panic(err)
	}

	return serialized.Bytes()
}

var MaxNonce = math.MaxUint32

func (block *Block) Hash() []byte {
	hash := sha256.Sum256(block.Serialize())
	return hash[:]
}

func (block *Block) ProofOfWork(difficulty int) error {

	if difficulty > 0 {

		nonce := 0
		isOdd := difficulty % 2

	PoW:
		for nonce < MaxNonce {
			block.Header.Nonce = uint32(nonce)
			hash := block.Hash()

			target := (difficulty - 1) / 2
			for i := 0; i < target; i++ {

				if hash[i] != 0 {
					nonce++
					goto PoW
				}

			}
			if (hash[target] > 15 && isOdd == 1) || (hash[target] != 0 && isOdd == 0) {
				nonce++
				goto PoW
			}

			block.Header.Hash = hash
			return nil
		}
	}

	return errors.New("Fail")
}

func NewGenesisBlock() *Block {

	block := Block{
		Header: BlockHeader{
			index:         0,
			PrevBlockHash: []byte{0},
			MerkleRoot:    []byte{0},
			Timestamp:     time.Date(2018, time.March, 3, 21, 0, 0, 0, time.UTC).String(),
			Difficulty:    1,
			MiningReward:  0,
			Nonce:         0,
			Hash:          []byte{0},
		},
		Transactions: []Transaction{{From: []byte{0}, To: []byte{0}, Value: 0}},
	}

	block.Header.Hash = block.Hash()

	return &block
}

func NewBlock(prevblockheader BlockHeader, transactions []Transaction) *Block {

	merkletree, err := NewTree(transactions)
	if err != nil {
		log.Panic(err)
	}

	difficulty := 2
	miningreward := 100

	block := Block{
		Header: BlockHeader{
			index:         prevblockheader.index + 1,
			PrevBlockHash: prevblockheader.Hash,
			MerkleRoot:    merkletree.merkleRoot,
			Timestamp:     time.Now().UTC().String(),
			Difficulty:    difficulty,
			MiningReward:  miningreward,
			Nonce:         0,
			Hash:          []byte{0},
		},
		Transactions: transactions,
	}

	err = block.ProofOfWork(difficulty)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
