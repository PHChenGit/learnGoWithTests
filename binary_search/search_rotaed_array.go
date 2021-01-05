package binary_search

func SearchRotatedArray(nums []int, target int) int  {
	lenNums := len(nums)

	if lenNums == 0 || (lenNums == 1 && nums[0] != target){
		return -1
	}

	if lenNums == 1 && nums[0] == target{
		return 0
	}

	left := 0
	right := lenNums-1
	var mid int

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid]{
			if nums[left] <= target && target <= nums[mid] {
				right = mid-1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid-1
			}
		}
	}

	return -1
}
