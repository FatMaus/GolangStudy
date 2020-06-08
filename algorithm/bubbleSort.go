package algorithm

// BubbleSort 冒泡排序，复杂度O(N^2)
func BubbleSort(numArr []int) {
	// 排好n-1个元素则全体排列完毕
	for i := 0; i < len(numArr)-1; i++ {
		// 循环n-1次的冒泡(升序)
		for n := 1; n < len(numArr); n++ {
			if numArr[n-1] > numArr[n] {
				var temp = numArr[n-1]
				numArr[n-1] = numArr[n]
				numArr[n] = temp
			} else {
				continue
			}
		}
	}
}
