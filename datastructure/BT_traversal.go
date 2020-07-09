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

// PreOrderTrav 前序遍历，将遍历结果装入切片返回
func PreOrderTrav(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int = make([]int, 0)
	var stack []*TreeNode = make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.value)
			stack = append(stack, root)
			root = root.leftNode
		}
		var node *TreeNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.rightNode
	}
	return result
}

// InOrderTrav 中序遍历，将遍历结果装入切片返回
func InOrderTrav(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		result []int       = make([]int, 0)
		stack  []*TreeNode = make([]*TreeNode, 0)
		node   *TreeNode
	)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.leftNode
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.value)
		root = node.rightNode
	}
	return result
}

// PostOrderTrav 后序遍历，将遍历结果装入切片返回
func PostOrderTrav(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		result    []int       = make([]int, 0)
		stack     []*TreeNode = make([]*TreeNode, 0)
		node      *TreeNode
		lastVisit *TreeNode
	)
	for root != nil || len(stack)-1 > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.leftNode
		}
		node = stack[len(stack)-1]
		if node.rightNode == nil || node.rightNode == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.value)
			lastVisit = node
		} else {
			root = node.rightNode
		}
	}
	return result
}

// PrevOrder 利用深度搜索进行前序遍历
func PrevOrder(root *TreeNode) []int {
	var result []int = make([]int, 0)
	Dfs(root, &result)
	return result
}

// Dfs 自上而下深度搜索
func Dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.value)
	Dfs(root.leftNode, result)
	Dfs(root.rightNode, result)
}

// PrevOrderT 利用自下而上分治法进行深度搜索，实现前序遍历
func PrevOrderT(root *TreeNode) []int {
	var result []int = divideDfs(root)
	return result
}

// divideDfs 分治法自下而上深度搜索
func divideDfs(root *TreeNode) []int {
	var (
		result []int = make([]int, 0)
		left   []int
		right  []int
	)
	if root == nil {
		return result
	}
	left = divideDfs(root.leftNode)
	right = divideDfs(root.rightNode)
	result = append(result, root.value)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// Bfs 广度优先遍历
func Bfs(root *TreeNode) [][]int {
	var (
		result   [][]int     = make([][]int, 0)
		queue    []*TreeNode = make([]*TreeNode, 0)
		level    *TreeNode
		floorNum int
	)
	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		var slice []int = make([]int, 0)
		floorNum = len(queue)
		for i := 0; i < floorNum; i++ {
			level = queue[0]
			queue = queue[1:]
			slice = append(slice, level.value)
			if level.leftNode != nil {
				queue = append(queue, level.leftNode)
			}
			if level.rightNode != nil {
				queue = append(queue, level.rightNode)
			}
		}
		result = append(result, slice)
	}
	return result
}

// MaxDepth 求二叉树最大深度，分治法+递归
func MaxDepth(root *TreeNode) int {
	var (
		ret   int
		left  int
		right int
	)
	if root == nil {
		ret = 0
	} else {
		left = MaxDepth(root.leftNode)
		right = MaxDepth(root.rightNode)
		switch {
		case left > right:
			ret = left + 1
		case right >= left:
			ret = right + 1
		}
	}
	return ret
}
