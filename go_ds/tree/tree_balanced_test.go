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

	lde := Height(root.left)
	rde := Height(root.right)
	flag := false
	if math.Abs(float64(lde-rde)) <= 1 {
		flag = true
	}
	return flag && isBalanced(root.left) && isBalanced(root.right)
}

func TestBalanced(t *testing.T) {
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
	n3.left = NewNode(41)
	n3.right = NewNode(42)

	fmt.Println(isBalanced(root))
}
