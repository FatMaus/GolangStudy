package algorithm

// RotateArray 求旋转数组中的最小元素，要求时间复杂度O(logN)
func RotateArray(numArray []int) int {
	var (
		start int = 0
		end   int = len(numArray) - 1
		mid   int = (end - start) / 2
	)
	for end-start > 1 {
		if numArray[mid] > numArray[end] {
			start = mid
			mid = start + (end-start)/2
		} else {
			end = mid
			mid = (end - start) / 2
		}
	}
	if numArray[start] < numArray[end] {
		return numArray[start]
	}
	return numArray[end]
}
