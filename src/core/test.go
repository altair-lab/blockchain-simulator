package core

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
	blockh.Timestamp = time.Date(2018, time.March, 3, 21, 0, 0, 123, time.UTC).String()
	blockh.Timestamp = time.Now().UTC().String()

	block := Block{}
	block.Header = blockh
	block.Transactions = tx

	blockchain := NewBlockChain()
	blockchain.CreateChain()
	blockchain.NewChain(tx)
	//NewChain(tx).Push()
	blockchain.Push(&block)

}
