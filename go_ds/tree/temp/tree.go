package temp

import (
	"fmt"
	"math"
)

// 判断一棵二叉树是不是满二叉树(树的节点层数h不超过16层)
type Node struct {
	val   int
	left  *Node
	right *Node
}

//二叉树深度
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lhight := maxDepth(root.left)
	rhight := maxDepth(root.right)
	if lhight > rhight {
		return lhight + 1
	} else {
		return rhight + 1
	}
}

/**
是否平衡二叉树
*/
func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}

	lde := maxDepth(root.left)
	rde := maxDepth(root.right)
	flag := false
	if math.Abs(float64(lde-rde)) <= 1 {
		flag = true
	}
	return flag && isBalanced(root.left) && isBalanced(root.right)
}

/**
满二叉树
*/
func isFull(root *Node) bool {
	if root == nil {
		return true
	}
	lheight := maxDepth(root.left)
	rheight := maxDepth(root.right)
	return isFull(root.left) && isFull(root.right) && (lheight == rheight)
}

/**
完全二叉树
*/
func isCompleted(root *Node) bool {
	if root == nil {
		return true
	}

	lheight := maxDepth(root.left)
	rheight := maxDepth(root.right)
	//完全二叉树
	flag := false
	if math.Abs(float64(lheight-rheight)) == 1 {
		flag = true
	} else {
		flag = false
	}
	return flag && isFull(root.left) && isFull(root.right)
}

// 中序遍历 左 -> 中 -> 右。
func MidOrderTraversal(tree *Node) {
	if tree == nil {
		return
	}

	MidOrderTraversal(tree.left)
	fmt.Printf(" %d -> ", tree.val)
	MidOrderTraversal(tree.right)
}

// 后序遍历   左 -> 右 -> 中。
func PostOrderTraversal(tree *Node) {
	if tree == nil {
		return
	}

	PostOrderTraversal(tree.left)
	PostOrderTraversal(tree.right)
	fmt.Printf(" %d -> ", tree.val)
}

// 前序遍历 遍历顺序
//  中 -> 左 -> 右。
func PreOrderTraversal(tree *Node) {
	if tree == nil {
		return
	}
	fmt.Printf(" %d -> ", tree.val)
	PreOrderTraversal(tree.left)
	PreOrderTraversal(tree.right)
}

// 按层遍历
func LevelOrderTraversal(tree *Node) {
	if tree == nil {
		return
	}

	// 采用队列实现
	queue := make([]*Node, 0)
	// queue push
	queue = append(queue, tree)
	for len(queue) > 0 {
		tree = queue[0]
		fmt.Printf(" %d -> ", tree.val)

		// queue pop
		queue = queue[1:]

		if tree.left != nil {
			queue = append(queue, tree.left)
		}
		if tree.right != nil {
			queue = append(queue, tree.right)
		}
	}
}
