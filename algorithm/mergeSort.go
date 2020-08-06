package algorithm

// MergeSort 归并排序，从中拆开二分，再分别递归至只有一个元素
func MergeSort(numArr []int, start int, end int) {
	// 拆至只有一个元素时停止
	if end-start < 1 {
		return
	}
	// 不断从中拆分递归
	var middle int = (start + end) / 2
	MergeSort(numArr, start, middle)
	MergeSort(numArr, middle+1, end)
	// 拆分后再排序
	merge(numArr, start, middle, end)
}

// merge 排序函数，主要比较中位元素和两侧元素
func merge(numArr []int, left int, mid int, right int) {
	// 拷贝数组进入一个辅助数组，以免影响原数组的元素
	// 设置辅助索引，进行排序
	var (
		leftStart  int   = left
		rightStart int   = mid + 1
		helpArr    []int = make([]int, len(numArr), len(numArr))
	)
	copy(helpArr, numArr)
	for i := left; i <= right; i++ {
		if leftStart > mid {
			// 此时表明偶数个元素的数组已被拆至最简
			numArr[i] = helpArr[rightStart]
			rightStart++
		} else if rightStart > right {
			// 表明奇数个元素的数组已被拆至最简
			numArr[i] = helpArr[leftStart]
			leftStart++
		} else if helpArr[leftStart] < helpArr[rightStart] {
			// 进行排序，符合规矩则维持原样，否则调换位置
			numArr[i] = helpArr[leftStart]
			leftStart++
		} else {
			numArr[i] = helpArr[rightStart]
			rightStart++
		}
	}
}
