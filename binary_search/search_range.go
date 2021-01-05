package binary_search

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	if len(nums) == 0 {
		return res
	}

	left := 0
	right := len(nums)-1
	var mid int

	for left < right {
		mid = left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] != target {
		return res
	}

	res[0] = left
	right = len(nums)-1

	for left < right {
		mid = left + (right-left)/2+1

		if nums[mid] > target {
			right = mid -1
		} else {
			left = mid
		}
	}

	res[1] = right
	return res
}
