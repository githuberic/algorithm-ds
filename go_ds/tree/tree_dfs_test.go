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
		if cur.Left != nil {
			s.Push(cur.Left)
		}
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		reS.Push(cur)
	}

	var res []interface{}
	for !reS.Empty() {
		res = append(res, reS.Pop().(*Node).Val)
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
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	n3.Left = NewNode(41)
	n3.Right = NewNode(42)

	values := root.DFS()
	fmt.Printf("DFS Order Value %v\n", values)
}
