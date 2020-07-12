package binarytree

func traversalLevelToLevel(root *TreeNode) [][]int {
	var (
		ret         [][]int     = make([][]int, 0)
		nodeQueue   []*TreeNode = make([]*TreeNode, 0)
		level       *TreeNode
		floorSlice  []int
		floorLength int
	)
	if root == nil {
		return ret
	}
	nodeQueue = append(nodeQueue, root)
	for len(nodeQueue) > 0 {
		floorLength = len(nodeQueue)
		floorSlice = make([]int, 0) // 每次初始化避免冗余输出
		for i := 0; i < floorLength; i++ {
			level = nodeQueue[0]
			nodeQueue = nodeQueue[1:]
			floorSlice = append(floorSlice, level.value)
			if level.leftNode != nil {
				nodeQueue = append(nodeQueue, level.leftNode)
			}
			if level.rightNode != nil {
				nodeQueue = append(nodeQueue, level.rightNode)
			}
		}
		ret = append(ret, floorSlice)
	}
	return ret
}

func traversalLevelToLevelReverse(root *TreeNode) [][]int {
	var ret [][]int = traversalLevelToLevel(root)
	reverse(ret)
	return ret
}

func reverse(slice [][]int) {
	var (
		temp []int
		j    int = len(slice) - 1
	)
	for i := 0; i < j; i++ {
		temp = slice[i]
		slice[i] = slice[j]
		slice[j] = temp
		j--
	}
}
