package array_and_slice

func Rotate(nums []int, k int) []int  {
	//numLength := len(nums)
	//for i := 0; i < k; i++ {
	//	for j := 0; j < numLength; j++ {
	//		nums[j], nums[numLength-1] = nums[numLength-1], nums[j]
	//	}
	//}
	//return nums
	numLength := len(nums)
	rotateArr := make([]int, numLength)
	for i := 0; i < numLength; i++ {
		rotateArr[(i+k)%numLength] = nums[i]
	}
	copy(nums, rotateArr)

	return rotateArr
}
