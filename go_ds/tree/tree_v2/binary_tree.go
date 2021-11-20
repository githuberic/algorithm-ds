package tree_v2

import (
	"algorithm-ds/go_ds/tree"
	"fmt"
)

type BinaryTree struct {
	Root *tree.Node
}

func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{tree.NewNode(v)}
}

/**
前序遍历
*/
func (this *BinaryTree) PreOrderTraverse() {
	p := this.Root
	s := tree.NewArrayStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			fmt.Printf("%+v ", p.Val)
			s.Push(p)
			p = p.Left
		} else {
			p = s.Pop().(*tree.Node).Right
		}
	}
}

/**
中序遍历
 */
func (this *BinaryTree) InOrderTraverse() {
	p := this.Root
	s := tree.NewArrayStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			s.Push(p)
			p = p.Left
		} else {
			tmp := s.Pop().(*tree.Node)
			fmt.Printf("%+v ", tmp.Val)
			p = tmp.Right
		}
	}
}

/***
后续遍历
 */
func (this *BinaryTree) PostOrderTraverse() {
	s1 := tree.NewArrayStack()
	s2 := tree.NewArrayStack()
	s1.Push(this.Root)
	for !s1.IsEmpty() {
		p := s1.Pop().(*tree.Node)
		s2.Push(p)
		if nil != p.Left {
			s1.Push(p.Left)
		}
		if nil != p.Right {
			s1.Push(p.Right)
		}
	}

	for !s2.IsEmpty() {
		fmt.Printf("%+v ", s2.Pop().(*tree.Node).Val)
	}
}
