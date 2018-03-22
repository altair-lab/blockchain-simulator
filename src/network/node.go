package main

import (
	"time"
	"fmt"
)

type node struct {
	performance int
}

func newNode(p int) *node {
	n := node{}
	n.performance = p
	return &n
}

func (n node) mining() {
	// consensus algorithm

	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	go func() {
		<-timer.C
		fmt.Println("mining")
	}()
	stop := timer.Stop()
	if stop {
		fmt.Println("timer stopped")
	}
}

