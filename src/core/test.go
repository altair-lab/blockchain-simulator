package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func Hash(block string) []byte {
	hash := sha256.Sum256([]byte(block))
	fmt.Printf("%x\n", hash)
	return hash[:]
}

func main() {
	tx := []Transaction{}
	tx = append(tx, Transaction{})
	tx = append(tx, Transaction{})
	tx = append(tx, Transaction{})
	tx[0].From = Hash("aaaa")

	blockh := BlockHeader{}

	tree, err := NewTree(tx)
	if err != nil {

	}
	blockh.MerkleRoot = tree.merkleRoot
	blockh.Timestamp = time.Now().Unix()

	block := Block{}
	block.Header = blockh
	block.Transactions = tx

	newblock := CreateBlock(blockh, tx)
	fmt.Printf("%d", newblock.Header.Nonce)

}
