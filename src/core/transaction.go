package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Transaction struct {
	From  []byte
	To    []byte
	Value uint64
}

func (tx Transaction) Serialize() []byte {
	var serialized bytes.Buffer
	encoder := gob.NewEncoder(&serialized)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return serialized.Bytes()
}

func (tx Transaction) Hash() []byte {
	hash := sha256.Sum256(tx.Serialize())
	return hash[:]
}
