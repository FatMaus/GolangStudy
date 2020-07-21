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

func findMinInRepeatArr(nums []int) int {
	// 思路：跳过重复元素，mid值和end值比较，分为两种情况进行处理
	if len(nums) == 0 {
		return -1
	}
	var (
		left  int = 0
		right int = len(nums) - 1
		mid   int
	)
	for left+1 < right {
		// 循环去重，直至去净
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		mid = left + (right-left)/2
		// 中间元素和最后一个元素比较（判断中间点落在左边上升区，还是右边上升区）
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[left] > nums[right] {
		return nums[right]
	}
	return nums[left]
}

func searchInRotatedArr(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	var (
		left  int = 0
		right int = len(nums) - 1
		mid   int
	)
	if target == nums[left] {
		return left
	} else if target == nums[right] {
		return right
	}
	for right-left > 1 {
		mid = (left + right) / 2
		if nums[mid] > target {
			if nums[mid] > nums[right] && target < nums[right] {
				left = mid
			} else {
				right = mid
			}
		} else if nums[mid] < target {
			if nums[mid] < nums[left] && target > nums[right] {
				right = mid
			} else {
				left = mid
			}
		} else {
			return mid
		}
	}
	return -1
}
