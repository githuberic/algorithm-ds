package tree

import "fmt"

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{NewNode(v)}
}

func (this *BinaryTree) InOrderTraverse() {
	p := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			s.Push(p)
			p = p.Left
		} else {
			tmp := s.Pop().(*Node)
			fmt.Printf("%+v ", tmp.Val)
			p = tmp.Right
		}
	}
}

func (this *BinaryTree) PreOrderTraverse() {
	p := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			fmt.Printf("%+v ", p.Val)
			s.Push(p)
			p = p.Left
		} else {
			p = s.Pop().(*Node).Right
		}
	}
}

func (this *BinaryTree) PostOrderTraverse() {
	s1 := NewArrayStack()
	s2 := NewArrayStack()
	s1.Push(this.root)
	for !s1.IsEmpty() {
		p := s1.Pop().(*Node)
		s2.Push(p)
		if nil != p.Left {
			s1.Push(p.Left)
		}
		if nil != p.Right {
			s1.Push(p.Right)
		}
	}

	for !s2.IsEmpty() {
		fmt.Printf("%+v ", s2.Pop().(*Node).Val)
	}
}

//use one stack, pre cursor to traverse from post order
func (this *BinaryTree) PostOrderTraverse2() {
	r := this.root
	s := NewArrayStack()

	//point to last visit node
	var pre *Node

	s.Push(r)

	for !s.IsEmpty() {
		r = s.Top().(*Node)
		if (r.Left == nil && r.Right == nil) || (pre != nil && (pre == r.Left || pre == r.Right)) {
			fmt.Printf("%+v ", r.Val)
			s.Pop()
			pre = r
		} else {
			if r.Right != nil {
				s.Push(r.Right)
			}
			if r.Left != nil {
				s.Push(r.Left)
			}
		}
	}
}