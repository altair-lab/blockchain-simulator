package core

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

//MerkleTree is the container for the tree. It holds a pointer to the root of the tree,
//a list of pointers to the leaf nodes, and the merkle root.
type MerkleTree struct {
	Root       *Node
	merkleRoot []byte
	Leafs      []*Node
}

//Node represents a node, root, or leaf in the tree. It stores pointers to its immediate
//relationships, a hash, the content stored if it is a leaf, and other metadata.
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	leaf   bool
	dup    bool
	Hash   []byte
	Tx     Transaction
}

//verifyNode walks down the tree until hitting a leaf, calculating the hash at each level
//and returning the resulting hash of Node n.
func (n *Node) verifyNode() []byte {
	if n.leaf {
		return n.Tx.Hash()
	}
	h := sha256.New()
	h.Write(append(n.Left.verifyNode(), n.Right.verifyNode()...))
	return h.Sum(nil)
}

//calculateNodeHash is a helper function that calculates the hash of the node.
func (n *Node) calculateNodeHash() []byte {
	if n.leaf {
		return n.Tx.Hash()
	}
	h := sha256.New()
	h.Write(append(n.Left.Hash, n.Right.Hash...))
	return h.Sum(nil)
}

//NewTree creates a new Merkle Tree using the content cs.
func NewTree(tx []Transaction) (*MerkleTree, error) {
	root, leafs, err := buildWithTx(tx)
	if err != nil {
		return nil, err
	}
	t := &MerkleTree{
		Root:       root,
		merkleRoot: root.Hash,
		Leafs:      leafs,
	}
	return t, nil
}

//buildWithContent is a helper function that for a given set of Contents, generates a
//corresponding tree and returns the root node, a list of leaf nodes, and a possible error.
//Returns an error if cs contains no Contents.
func buildWithTx(tx []Transaction) (*Node, []*Node, error) {
	if len(tx) == 0 {
		return nil, nil, errors.New("Error: cannot construct tree with no content.")
	}
	var leafs []*Node
	for _, t := range tx {
		leafs = append(leafs, &Node{
			Hash: t.Hash(),
			Tx:   t,
			leaf: true,
		})
	}
	if len(leafs)%2 == 1 {
		duplicate := &Node{
			Hash: leafs[len(leafs)-1].Hash,
			Tx:   leafs[len(leafs)-1].Tx,
			leaf: true,
			dup:  true,
		}
		leafs = append(leafs, duplicate)
	}
	root := buildIntermediate(leafs)
	return root, leafs, nil
}

//buildIntermediate is a helper function that for a given list of leaf nodes, constructs
//the intermediate and root levels of the tree. Returns the resulting root node of the tree.
func buildIntermediate(nl []*Node) *Node {
	var nodes []*Node
	for i := 0; i < len(nl); i += 2 {
		h := sha256.New()
		var left, right int = i, i + 1
		if i+1 == len(nl) {
			right = i
		}
		chash := append(nl[left].Hash, nl[right].Hash...)
		h.Write(chash)
		n := &Node{
			Left:  nl[left],
			Right: nl[right],
			Hash:  h.Sum(nil),
		}
		nodes = append(nodes, n)
		nl[left].Parent = n
		nl[right].Parent = n
		if len(nl) == 2 {
			return n
		}
	}
	return buildIntermediate(nodes)
}

//String returns a string representation of the tree. Only leaf nodes are included
//in the output.
func (m *MerkleTree) String() string {
	s := ""
	for _, l := range m.Leafs {
		s += fmt.Sprint(l)
		s += "\n"
	}
	return s
}
