package temp

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := &Node{1, nil, nil}
	n1 := &Node{2, nil, nil}
	n2 := &Node{3, nil, nil}
	n3 := &Node{4, nil, nil}
	n4 := &Node{5, nil, nil}
	//n5 := &Node{6, nil, nil}
	//n6 := &Node{7, nil, nil}

	root.left = n1
	root.right = n2
	n1.left = n3
	n1.right = n4
	//n2.left = n5
	//n2.right = n6
	fmt.Println("前------------")
	PreOrderTraversal(root)
	fmt.Println()
	fmt.Println("中------------")
	MidOrderTraversal(root)
	fmt.Println()
	fmt.Println("后------------")
	PostOrderTraversal(root)
	fmt.Println()
	fmt.Println("层------------")
	LevelOrderTraversal(root)
	fmt.Println("------------")
	fmt.Println(isFull(root))
}
