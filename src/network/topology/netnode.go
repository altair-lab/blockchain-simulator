package main

// Maximum number of nodes is specified.
var nodeCnt uint

// NetNodeID - Each node has its unique ID.
type NetNodeID uint

// NodeType - Node type specifies whether this node is static/dynamic full/light/archive etc.
type NodeType uint8

// NetNode - Network node type.
type NetNode struct {
	nodeType NodeType // TODO
	id       NetNodeID
	nodePtr  *Node   // TODO
	subnet   *Subnet // TODO
}

func (nNet NetNode) newNode() *NetNode {

}

func newNode() NetNode {

}
