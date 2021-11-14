package tree

import (
	"fmt"
	"testing"
)

// 按层遍历
func LevelOrderTraversal(tree *Node) []interface{} {
	if tree == nil {
		return nil
	}

	var res []interface{}

	// 采用队列实现
	queue := make([]*Node, 0)
	// queue push
	queue = append(queue, tree)
	for len(queue) > 0 {
		tree = queue[0]
		res = append(res, tree.data)

		// queue pop
		queue = queue[1:]
		if tree.left != nil {
			queue = append(queue, tree.left)
		}
		if tree.right != nil {
			queue = append(queue, tree.right)
		}
	}
	return res
}

/**
验证是否满二叉树
*/
func TestBFS(t *testing.T) {
	root := NewNode(1)
	n1 := NewNode(2)
	n2 := NewNode(3)
	n3 := NewNode(4)
	n4 := NewNode(5)
	root.left = n1
	root.right = n2
	n1.left = n3
	n1.right = n4
	n4.left = NewNode(41)
	n4.right = NewNode(42)

	values := LevelOrderTraversal(root)
	fmt.Printf("BFS Order Value %v\n", values)
}
