package algorithm

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var (
		left  int = 0
		right int = len(nums) - 1
		mid   int = (left + right) / 2
	)
	if right == 0 {
		return nums[0]
	}
	for right >= left {
		mid = (left + right) / 2
		switch {
		case nums[mid] > nums[right]:
			left = mid + 1
		case mid > 0 && nums[mid] < nums[mid-1]:
			return nums[mid]
		default:
			right = mid - 1
		}
	}
	return nums[mid]
}

func findMinInRepeatArr(nums []int) {

}
