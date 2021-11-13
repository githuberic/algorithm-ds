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
	res = append(res, root.data)
	res = append(res, preOrderTraversal(root.left)...)
	res = append(res, preOrderTraversal(root.right)...)

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

	res := inOrderTraversal(root.left)
	res = append(res, root.data)
	res = append(res, inOrderTraversal(root.right)...)

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
	if root.left != nil {
		lRes := postOrderTraversal(root.left)
		if len(lRes) > 0 {
			res = append(res, lRes...)
		}
	}
	if root.right != nil {
		rRes := postOrderTraversal(root.right)
		if len(rRes) > 0 {
			res = append(res, rRes...)
		}
	}
	res = append(res, root.data)
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
	root.left = n1
	root.right = n2
	n1.left = n3
	n1.right = n4

	values := preOrderTraversal(root)
	fmt.Printf("Pre Order Value %v\n", values)

	inOrderValues := inOrderTraversal(root)
	fmt.Printf("In order Value %v\n", inOrderValues)

	postOrderValues := postOrderTraversal(root)
	fmt.Printf("Post order Value %v\n", postOrderValues)
}
