package tree

import (
	"fmt"
	"testing"
)

//二叉树深度
func Height(root *Node) int {
	if root == nil {
		return 0
	}

	lHeight := Height(root.Left)
	rHeight := Height(root.Right)
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
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	fmt.Println(Height(root))
}
