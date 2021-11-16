package tree

import (
	"fmt"
	"testing"
)

/**
Tree深度优先遍历
*/
func (tree *Node) DFS() []interface{} {
	if tree == nil {
		return nil
	}

	reS := NewStack()
	s := NewStack()
	s.Push(tree)
	for !s.Empty() {
		cur := s.Pop().(*Node)
		if cur.left != nil {
			s.Push(cur.left)
		}
		if cur.right != nil {
			s.Push(cur.right)
		}
		reS.Push(cur)
	}

	var res []interface{}
	for !reS.Empty() {
		res = append(res, reS.Pop().(*Node).data)
	}

	return res
}

/**
Tree深度遍历测试
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

	n3.left = NewNode(41)
	n3.right = NewNode(42)

	values := root.DFS()
	fmt.Printf("DFS Order Value %v\n", values)
}
