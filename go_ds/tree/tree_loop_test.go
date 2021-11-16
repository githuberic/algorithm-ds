package tree

import (
	"fmt"
	"testing"
)

/**
前序遍历

中 -> 左 -> 右
*/
func preOrderTraversal(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	var res []interface{}
	res = append(res, root.Val)
	res = append(res, preOrderTraversal(root.Left)...)
	res = append(res, preOrderTraversal(root.Right)...)

	return res
}

/**
中序遍历
 左 -> 中 -> 右
*/
func inOrderTraversal(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	res := inOrderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inOrderTraversal(root.Right)...)

	return res
}

/**
后续遍历
左 -> 右 -> 中
*/
func postOrderTraversal(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	var res []interface{}
	if root.Left != nil {
		lRes := postOrderTraversal(root.Left)
		if len(lRes) > 0 {
			res = append(res, lRes...)
		}
	}
	if root.Right != nil {
		rRes := postOrderTraversal(root.Right)
		if len(rRes) > 0 {
			res = append(res, rRes...)
		}
	}
	res = append(res, root.Val)
	return res
}

/**
验证是否满二叉树
*/
func TestLoop(t *testing.T) {
	root := NewNode(1)
	n1 := NewNode(2)
	n2 := NewNode(3)
	n3 := NewNode(4)
	n4 := NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	values := preOrderTraversal(root)
	fmt.Printf("Pre Order Value %v\n", values)

	inOrderValues := inOrderTraversal(root)
	fmt.Printf("In order Value %v\n", inOrderValues)

	postOrderValues := postOrderTraversal(root)
	fmt.Printf("Post order Value %v\n", postOrderValues)
}
