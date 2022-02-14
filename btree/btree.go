package btree

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Operation interface {
	CreateBTree(node *Node)

	// PreOrderTraverse 前序遍历
	PreOrderTraverse(node *Node)
	// InOrderTraverse 中序遍历
	InOrderTraverse(node *Node)
	// PostOrderTraverse 后序遍历
	PostOrderTraverse(node *Node)

	Search(value int)
	Insert(value int)
}

func InitBinaryTree(root *Node) *Node {
	l := &Node{}
	r := &Node{}

	root.left = l.NewBinaryTreeNode(2)
	root.right = r.NewBinaryTreeNode(3)

	l2 := &Node{}
	ll2 := &Node{}

	root.left.left = l2.NewBinaryTreeNode(4)
	root.right.right = ll2.NewBinaryTreeNode(5)
	return root
}

func (n *Node) NewBinaryTreeNode(value int) *Node {
	n.value = value
	n.left = nil
	n.right = nil
	return n
}

// GetNodeNum 计算二叉树的节点个数
func GetNodeNum(root *Node) int {
	if root != nil {
		return GetNodeNum(root.left) + GetNodeNum(root.right) + 1
	}
	return 0
}

// GetDegress 获取二叉树的深度
func GetDegress(root *Node) int {
	if root == nil {
		return 0
	}
	var maxDegress = 0
	if GetDegress(root.left) > GetDegress(root.right) {
		maxDegress = GetDegress(root.left)
	} else {
		maxDegress = GetDegress(root.right)
	}

	return maxDegress + 1
}

// PreOrder 前序遍历
func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println("%d->", root.value)
	PreOrder(root.left)
	PreOrder(root.right)
}

func InOrder(root *Node) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Printf("%d->", root.value)
	InOrder(root.right)
}

// PostOrder 后续遍历
func PostOrder(root *Node) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d->", root.value)
}

// GetKthNum  获取K层节点个数
func GetKthNum(root *Node, k int) int {
	if root == nil {
		return 0
	}

	if k == 1 {
		return 1
	}

	return GetKthNum(root.left, k-1) + GetKthNum(root.right, k-1)
}

func GetLeavNum(root *Node) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1
	}

	return GetLeavNum(root.left) + GetLeavNum(root.right)
}

// IsBalanced 判断是否是平衡二叉树
func IsBalanced(root *Node) bool {
	if root == nil {
		return true
	}

	lde := GetDegress(root.left)
	rde := GetDegress(root.right)

	var flag = false
	if math.Abs(float64(lde-rde)) <= 1 {
		flag = true
	} else {
		flag = false
	}

	return flag && IsBalanced(root.left) && IsBalanced(root.right)
}

