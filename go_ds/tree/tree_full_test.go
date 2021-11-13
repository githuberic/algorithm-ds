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

	lHeight := Height(root.left)
	rHeight := Height(root.right)
	return isFull(root.left) && isFull(root.right) && (lHeight == rHeight)
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
	root.left = n1
	root.right = n2
	n1.left = n3
	n1.right = n4

	n2.left = NewNode(31)
	n2.right = NewNode(32)

	fmt.Println(isFull(root))
}
