package binarytree

// TreeNode node for binary tree
type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// PathCounter to count the maximum path sum
type PathCounter struct {
	singleMaxPath int
	totalMaxPath  int
}

// 路径经过的节点的值的和即为路径和
func binaryTreeMaxPathSum(root *TreeNode) int {
	var result *PathCounter
	result = PathCount(root)
	return result.totalMaxPath
}

// PathCount find maximum path
func PathCount(root *TreeNode) *PathCounter {
	var (
		ret     *PathCounter
		left    *PathCounter
		right   *PathCounter
		maxPath int
	)
	ret = &PathCounter{
		singleMaxPath: 0,
		totalMaxPath:  -(1 << 31),
	}
	if root == nil {
		return ret
	}
	left = PathCount(root.leftNode)
	right = PathCount(root.rightNode)
	maxPath = compare(left.totalMaxPath, right.totalMaxPath)
	if left.singleMaxPath > right.singleMaxPath {
		ret.singleMaxPath = compare(left.singleMaxPath+root.value, 0)
	} else {
		ret.singleMaxPath = compare(right.singleMaxPath+root.value, 0)
	}
	ret.totalMaxPath = compare(maxPath, left.singleMaxPath+right.singleMaxPath+root.value)
	return ret
}

func compare(num1, num2 int) int {
	var ret int
	if num1 > num2 {
		ret = num1
	} else {
		ret = num2
	}
	return ret
}

// CommonAncestor to find common ancestor for two TreeNodes
func CommonAncestor(root, p, q *TreeNode) *TreeNode {
	var (
		ret   *TreeNode
		left  *TreeNode
		right *TreeNode
	)
	if root == nil {
		return root
	} else if root == p || root == q {
		return root
	}
	left = CommonAncestor(root.leftNode, p, q)
	right = CommonAncestor(root.rightNode, p, q)
	if left != nil && right != nil {
		ret = root
	} else if left != nil {
		ret = left
	} else if right != nil {
		ret = right
	} else {
		ret = nil
	}
	return ret
}
