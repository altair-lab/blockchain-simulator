package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"log"
	"math"
)

type BlockHeader struct {
	PrevBlockHash []byte
	MerkleRoot    []byte
	Timestamp     int64
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
			PrevBlockHash: block.Header.PrevBlockHash,
			MerkleRoot:    block.Header.MerkleRoot,
			Timestamp:     block.Header.Timestamp,
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

func ProofOfWork(block *Block) error {

	nonce := 0
	difficulty := 3
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

	return errors.New("Fail")
}

func CreateBlock(blockheader BlockHeader, transactions []Transaction) *Block {
	block := &Block{blockheader, transactions}
	err := ProofOfWork(block)
	if err != nil {
		log.Panic(err)
	}
	return block
}
