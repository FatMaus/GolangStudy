package algorithm

// InsertSort 插入排序
func InsertSort(numArr []int) {
	// 挨个排序，只用排n-1个元素即可
	for i := 1; i < len(numArr); i++ {
		// 确定需要排列的元素，设定计数器
		var count int = i - 1
		var baseNum = numArr[i]
		// 从待排向前排查，放到第一个不大于该元素的元素旁
		for count >= 0 && numArr[count] > baseNum {
			numArr[count+1] = numArr[count]
			count--
		}
		numArr[count+1] = baseNum
	}
}
