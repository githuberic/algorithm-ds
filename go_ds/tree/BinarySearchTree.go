package tree

type BST struct {
	*BinaryTree
	//比对函数，0:v==nodeV,正数:v>nodeV,负数:v<nodeV
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	if nil == compareFunc {
		return nil
	}
	return &BST{BinaryTree: NewBinaryTree(rootV), compareFunc: compareFunc}
}

func (this *BST) Find(v interface{}) *Node {
	p := this.root
	for nil != p {
		compareResult := this.compareFunc(v, p.Val)
		if compareResult == 0 {
			return p
		} else if compareResult > 0 { //v > nodeV
			p = p.Right
		} else { //v < nodeV
			p = p.Left
		}
	}
	return nil
}

func (this *BST) Insert(v interface{}) bool {
	p := this.root
	for nil != p {
		compareResult := this.compareFunc(v, p.Val)
		if compareResult == 0 {
			return false
		} else if compareResult > 0 {
			if nil == p.Right {
				p.Right = NewNode(v)
				break
			}
			p = p.Right
		} else {
			if nil == p.Left {
				p.Left = NewNode(v)
				break
			}
			p = p.Left
		}
	}
	return true
}

func (this *BST) Delete(v interface{}) bool {
	var pp *Node = nil
	p := this.root
	deleteLeft := false
	for nil != p {
		compareResult := this.compareFunc(v, p.Val)
		if compareResult > 0 {
			pp = p
			p = p.Right
			deleteLeft = false
		} else if compareResult < 0 {
			pp = p
			p = p.Left
			deleteLeft = true
		} else {
			break
		}
	}

	if nil == p { //需要删除的节点不存在
		return false
	} else if nil == p.Left && nil == p.Right { //删除的是一个叶子节点
		if nil != pp {
			if deleteLeft {
				pp.Left = nil
			} else {
				pp.Right = nil
			}
		} else { //根节点
			this.root = nil
		}
	} else if nil != p.Right { //删除的是一个有右孩子，不一定有左孩子的节点
		//找到p节点右孩子的最小节点
		pq := p
		q := p.Right //向右走一步
		fromRight := true
		for nil != q.Left { //向左走到底
			pq = q
			q = q.Left
			fromRight = false
		}
		if fromRight {
			pq.Right = nil
		} else {
			pq.Left = nil
		}
		q.Left = p.Left
		q.Right = p.Right
		if nil == pp { //根节点被删除
			this.root = q
		} else {
			if deleteLeft {
				pq.Left = q
			} else {
				pq.Right = q
			}
		}
	} else { //删除的是一个只有左孩子的节点
		if nil != pp {
			if deleteLeft {
				pp.Left = p.Left
			} else {
				pp.Right = p.Left
			}
		} else {
			if deleteLeft {
				this.root = p.Left
			} else {
				this.root = p.Left
			}
		}
	}
	return true
}


func (this *BST) Min() *Node {
	p := this.root
	for nil != p.Left {
		p = p.Left
	}
	return p
}

func (this *BST) Max() *Node {
	p := this.root
	for nil != p.Right {
		p = p.Right
	}
	return p
}