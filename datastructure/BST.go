package datastructure

import "fmt"

// TreeNode 排序二叉树节点结构
type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// NewTreeNode 构造函数，构造一个新节点
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		value:     value,
		leftNode:  nil,
		rightNode: nil,
	}
}

// BST 排序二叉树结构
type BST struct {
	root *TreeNode
}

// NewBST 构造函数，构造一个新排序二叉树
func NewBST() *BST {
	return &BST{
		root: nil,
	}
}

// Get 找到与key相同节点的node，若无则返回空值
func (b *BST) Get(key int) *TreeNode {
	var cur = b.root
	for cur != nil {
		switch {
		case key < cur.value:
			cur = cur.leftNode
		case key > cur.value:
			cur = cur.rightNode
		case key == cur.value:
			return cur
		default:
			fmt.Println("something wrong")
			break
		}
	}
	return nil
}

// Insert 将新元素插入到合适的位置
func (b *BST) Insert(key int) {
	if b.root == nil {
		b.root = NewTreeNode(key)
		return
	}
	// 非空树进行迭代查找
	var cur = b.root
	var parent *TreeNode = nil
	for {
		parent = cur
		switch {
		case key < parent.value:
			cur = parent.leftNode
			if cur == nil {
				parent.leftNode = NewTreeNode(key)
				return
			}
		case key > parent.value:
			cur = parent.rightNode
			if cur == nil {
				parent.rightNode = NewTreeNode(key)
				return
			}
		default:
			fmt.Println("something going wrong")
			return
		}
	}
}

// Remove 移除指定的节点，若不存在该节点则返回false
func (b *BST) Remove(key int) bool {
	var cur = b.root
	var parent = cur
	var isLeftChild = false
	for cur != nil && cur.value != key {
		parent = cur
		if cur.value > key {
			isLeftChild = true
			cur = cur.leftNode
		} else {
			isLeftChild = false
			cur = cur.rightNode
		}
	}
	// 遍历完成，未找到则返回false
	if cur == nil {
		return false
	}
	// 节点存在，判断是leaf还是有子节点
	if cur.leftNode == nil && cur.rightNode == nil {
		switch {
		case cur == b.root:
			b.root = nil
		case isLeftChild:
			parent.leftNode = nil
		default:
			parent.rightNode = nil
		}
	} else if cur.rightNode == nil {
		switch {
		case cur == b.root:
			b.root = cur.leftNode
		case isLeftChild:
			parent.leftNode = cur.leftNode
		default:
			parent.rightNode = cur.leftNode
		}
	} else if cur.leftNode == nil {
		switch {
		case cur == b.root:
			b.root = cur.rightNode
		case isLeftChild:
			parent.leftNode = cur.rightNode
		default:
			parent.rightNode = cur.rightNode
		}
	} else {
		var replaceNode = getReplace(cur)
		switch {
		case cur == b.root:
			b.root = replaceNode
		case isLeftChild:
			parent.leftNode = replaceNode
		default:
			parent.rightNode = replaceNode
		}
		replaceNode.leftNode = cur.leftNode
		replaceNode.rightNode = cur.rightNode
	}
	return true
}

// 辅助函数，找出满足要求的替代节点。
// 左子树的最大值和右子树的最小值比较，较大者为替代节点
func getReplace(node *TreeNode) *TreeNode {
	var (
		left        = node.leftNode
		leftParent  = node
		right       = node.rightNode
		rightParent = node
	)
	for left.rightNode != nil {
		leftParent = left
		left = left.rightNode
	}
	for right.leftNode != nil {
		rightParent = right
		right = right.leftNode
	}
	if left.value > right.value {
		leftParent.rightNode = nil
		return left
	}
	rightParent.leftNode = nil
	return right
}

// Display 遍历打印二叉树，使用中序遍历
func Display(root *TreeNode) {
	if root == nil {
		return
	}
	Display(root.leftNode)
	fmt.Print(root.value, ", ")
	Display(root.rightNode)
}