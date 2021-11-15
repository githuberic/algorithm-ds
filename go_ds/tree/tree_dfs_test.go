package tree

import (
	"fmt"
	"testing"
)

// 按层遍历
func DFS(tree *Node) []interface{} {
	if tree == nil {
		return nil
	}

	s := NewStack()
	s.Push(tree)
	for !s.Empty() {
		cur := s.Pop().(*Node)

		if cur.left != nil {
			s.Push(tree.left)
		}

		if cur.right != nil {
			s.Push(tree.right)
		}

		s.Push(tree.data)
	}

	var res []interface{}
	for !s.Empty() {
		cur := s.Pop().(*Node)
		res = append(res, cur.data)
	}

	return res
}

/**
验证是否满二叉树
*/
func TestDFS(t *testing.T) {
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

	values := DFS(root)
	fmt.Printf("DFS Order Value %v\n", values)
}

// https://studygolang.com/articles/16314
