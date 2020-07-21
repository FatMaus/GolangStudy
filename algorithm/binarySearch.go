package algorithm

func binarySearch1(nums []int, target int) int {
	var (
		left  int = 0
		right int = len(nums) - 1
		mid   int
	)
	for right >= left {
		mid = (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func binarySearch2(nums []int, target int) int {
	var (
		left  int = 0
		right int = len(nums) - 1
		mid   int
	)
	for right-left > 1 {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	switch {
	case nums[right] == target:
		return right
	case nums[left] == target:
		return left
	default:
		return -1
	}
}
