package tree

import "testing"

func TestBinaryTree_InOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.Left = NewNode(3)
	binaryTree.root.Right = NewNode(4)
	binaryTree.root.Right.Left = NewNode(5)

	binaryTree.InOrderTraverse()
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.Left = NewNode(3)
	binaryTree.root.Right = NewNode(4)
	binaryTree.root.Right.Left = NewNode(5)

	binaryTree.PreOrderTraverse()
}

func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.Left = NewNode(3)
	binaryTree.root.Right = NewNode(4)
	binaryTree.root.Right.Left = NewNode(5)

	binaryTree.PostOrderTraverse()
}

func TestBinaryTree_PostOrderTraverse2(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.Left = NewNode(3)
	binaryTree.root.Right = NewNode(4)
	binaryTree.root.Right.Left = NewNode(5)

	binaryTree.PostOrderTraverse2()
}
