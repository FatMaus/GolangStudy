package algorithm

// HeapSort 堆排序，利用MaxHeap的性质进行排序
func HeapSort(nums []int) []int {
	// 构建maxHeap
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i, len(nums))
	}
	// 首尾交换并不断保持结构
	for i := len(nums) - 1; i >= 1; i-- {
		swap(nums, 0, i)
		sink(nums, 0, i)
	}
	return nums
}

func sink(nums []int, i int, length int) {
	var (
		left  int
		right int
		index int
	)
	for {
		// 设定左右索引
		left = i*2 + 1
		right = i*2 + 2
		index = i
		// 若左节点满足条件
		if left < length && nums[left] > nums[index] {
			index = left
		}
		// 若右节点满足条件
		if right < length && nums[right] > nums[index] {
			index = right
		}
		// 判断完后当前节点最大则工作完成
		if index == i {
			break
		}
		// 若产生了交换则继续下沉判断
		swap(nums, i, index)
		i = index
	}
}
