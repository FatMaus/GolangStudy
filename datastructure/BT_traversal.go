package datastructure

import "fmt"

// PreOrderTraversal 前序遍历
func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.value, ", ")
	PreOrderTraversal(root.leftNode)
	PreOrderTraversal(root.rightNode)
}

// PostOrderTraversal 后序遍历
func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.leftNode)
	PostOrderTraversal(root.rightNode)
	fmt.Print(root.value, ", ")
}

// InOrderTraversal 中序遍历，适用于排序二叉树
func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.leftNode)
	fmt.Print(root.value, ", ")
	InOrderTraversal(root.rightNode)
}
