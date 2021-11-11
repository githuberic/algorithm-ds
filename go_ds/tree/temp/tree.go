package temp

import (
	"fmt"
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
	return max(lhight, rhight)
}

func max(a, b int) int {
	if a > b {
		return a + 1
	}
	return b + 1
}

//判断
func isFull(root *Node) bool {
	if root == nil {
		return true
	}
	lheight := maxDepth(root.left)
	rheight := maxDepth(root.right)
	fmt.Println("l:", lheight, "r:", rheight)
	return isFull(root.left) && isFull(root.right) && (lheight == rheight)

	//完全二叉树
	//if math.Abs(float64(Height(root.left)-Height(root.right))) ==1  {
	//  flag = true
	//} else {
	//  flag = false
	//}
	//return flag && isFull(root.left) && isFull(root.right)
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