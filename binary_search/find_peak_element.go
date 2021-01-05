package binary_search

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums)-1
	var mid int

	for left < right {
		mid = left + (right - left) / 2

		if nums[mid] < nums[mid+1] {
			left = mid+1
		} else if nums[mid] <= nums[mid-1] {
			right = mid
		}
	}

	if left != len(nums)-1 && nums[left] > nums[left+1] {
		return left
	} else if right != 0 && nums[right] > nums[right-1]{
		return right
	}

	return mid
}