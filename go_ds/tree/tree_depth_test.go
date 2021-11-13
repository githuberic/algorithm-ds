package tree

import (
	"fmt"
	"testing"
)

//二叉树深度
func height(root *Node) int {
	if root == nil {
		return 0
	}

	lHeight := height(root.left)
	rHeight := height(root.right)
	if lHeight > rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}

func TestDepth(t *testing.T) {
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

	fmt.Println(height(root))
}
