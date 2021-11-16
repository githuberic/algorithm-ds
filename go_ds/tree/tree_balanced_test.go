package tree

import (
	"fmt"
	"math"
	"testing"
)

/**
是否平衡二叉树
*/
func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}

	lde := Height(root.Left)
	rde := Height(root.Right)
	flag := false
	if math.Abs(float64(lde-rde)) <= 1 {
		flag = true
	}
	return flag && isBalanced(root.Left) && isBalanced(root.Right)
}

func TestBalanced(t *testing.T) {
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
	n3.Left = NewNode(41)
	n3.Right = NewNode(42)

	fmt.Println(isBalanced(root))
}
