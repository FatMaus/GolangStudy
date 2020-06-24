package algorithm

// SelectionSort 选择排序
func SelectionSort(numArr []int) {
	for i := 0; i < len(numArr); i++ {
		var minNumIndex int = i
		for j := i; j < len(numArr); j++ {
			if numArr[minNumIndex] > numArr[j] {
				minNumIndex = j
			} else {
				continue
			}
		}
		var temp = numArr[minNumIndex]
		numArr[minNumIndex] = numArr[i]
		numArr[i] = temp
	}
}
