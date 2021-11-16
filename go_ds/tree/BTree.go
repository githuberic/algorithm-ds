package tree

type BTree struct {
	root *Node
}

func NewBTree(v interface{}) *BTree {
	return &BTree{NewNode(v)}
}
