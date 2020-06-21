package algorithm

func rotateArray(numArray []int) int {
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
