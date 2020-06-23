package algorithm

// QuickSort 快速排序函数，对一个数组进行快速排序
func QuickSort(numArr []int, left int, right int) {
	if left >= right {
		return
	}
	// 找到分治点
	partitionIndex := partition(numArr, left, right)
	// 二分法递归
	QuickSort(numArr, left, partitionIndex-1)
	QuickSort(numArr, partitionIndex+1, right)
}

// partition 通过基准元素快速分类，找到分治点以便分治
func partition(numArr []int, left int, right int) int {
	// 按兴趣取一个基准元素
	var baseIndex = right
	var pivot = numArr[baseIndex]
	var leftIndex int = left
	var rightIndex int = right - 1
	// 只用排n-1个元素即可，双向排序到中间会和
	for {
		for leftIndex < right && numArr[leftIndex] <= pivot {
			leftIndex++
		}
		for rightIndex >= left && numArr[rightIndex] > pivot {
			rightIndex--
		}
		// 会和则结束循环，未碰头则对不符合排序规则的元素进行调换分类
		if leftIndex > rightIndex {
			break
		} else {
			swap(numArr, leftIndex, rightIndex)
		}
	}
	// 基准元素放到合适的位置，返回分治点
	swap(numArr, leftIndex, baseIndex)
	return leftIndex
}

// swap 交换函数，用于交换两个元素的位置
func swap(numArr []int, left int, right int) {
	var temp = numArr[left]
	numArr[left] = numArr[right]
	numArr[right] = temp
}
