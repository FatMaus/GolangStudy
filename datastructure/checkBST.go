package datastructure

// 建立节点信息结构体
type isBSTNode struct {
	isBST    bool
	maxValue int
	minValue int
}

func isValidBST(root *TreeNode) bool {
	var ret *isBSTNode
	ret = checkBST(root)
	return ret.isBST
}

// 按照定义逐项对比，若某个子节点不是则一票否决
func checkBST(root *TreeNode) *isBSTNode {
	var (
		ret   *isBSTNode
		left  *isBSTNode
		right *isBSTNode
	)
	ret = &isBSTNode{}
	if root == nil {
		ret.isBST = true
		return ret
	}

	left = checkBST(root.leftNode)
	right = checkBST(root.rightNode)

	if !left.isBST || !right.isBST {
		ret.isBST = false
		return ret
	} else if left.maxValue >= root.value && root.leftNode != nil {
		ret.isBST = false
		return ret
	} else if root.value >= right.minValue && root.rightNode != nil {
		ret.isBST = false
		return ret
	}
	ret.isBST = true
	ret.minValue = root.value
	if root.leftNode != nil {
		ret.minValue = left.minValue
	}
	ret.maxValue = root.value
	if root.rightNode != nil {
		ret.maxValue = right.maxValue
	}
	return ret
}

// 也可使用中序遍历装入切片，再逐一比对
func isValidBST2(root *TreeNode) bool {
	var (
		checkSlice []int = make([]int, 0)
		isBST2     bool  = true
	)
	inOrderSlice(root, &checkSlice)
	for i := 1; i < len(checkSlice); i++ {
		if checkSlice[i-1] >= checkSlice[i] {
			isBST2 = false
			break
		}
	}
	return isBST2
}

func inOrderSlice(root *TreeNode, checkSlice *[]int) {
	if root == nil {
		return
	}
	inOrderSlice(root.leftNode, checkSlice)
	*checkSlice = append(*checkSlice, root.value)
	inOrderSlice(root.rightNode, checkSlice)
}
