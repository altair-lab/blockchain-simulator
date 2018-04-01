package network

import (
	"fmt"
	"time"
)

// Node.
type Node struct {
	performance int
}

func newNode(p int) *Node {
	n := Node{}
	n.performance = p
	return &n
}

func (n Node) mining() {
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
