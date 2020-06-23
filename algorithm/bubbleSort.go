package algorithm

// BubbleSort 冒泡排序，复杂度O(N^2)
func BubbleSort(numArr []int) {
	// 监测器，若无动作则表示排序完毕可直接结束，可减少重复步骤
	var action bool
	// 排好n-1个元素则全体排列完毕
	for i := 0; i < len(numArr)-1; i++ {
		action = false
		// 循环n-1次的冒泡(升序)
		for n := 1; n < len(numArr); n++ {
			if numArr[n-1] > numArr[n] {
				var temp = numArr[n-1]
				numArr[n-1] = numArr[n]
				numArr[n] = temp
				action = true
			} else {
				continue
			}
		}
		if !action {
			break
		}
	}
}
