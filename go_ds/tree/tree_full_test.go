package tree

import (
	"fmt"
	"testing"
)

/**
满二叉树
*/
func isFull(root *Node) bool {
	if root == nil {
		return true
	}

	lHeight := Height(root.Left)
	rHeight := Height(root.Right)
	return isFull(root.Left) && isFull(root.Right) && (lHeight == rHeight)
}

/**
验证是否满二叉树
*/
func TestFull(t *testing.T) {
	//root := &Node{1, nil, nil}
	root := NewNode(1)
	n1 := NewNode(2)
	n2 := NewNode(3)
	n3 := NewNode(4)
	n4 := NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	n2.Left = NewNode(31)
	n2.Right = NewNode(32)

	fmt.Println(isFull(root))
}
