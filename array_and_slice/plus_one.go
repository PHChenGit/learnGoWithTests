package array_and_slice


func PlusOne(digits []int) []int {
	plusOne := false
	tailIdx := len(digits) - 1
	for i := tailIdx; i >= 0; i-- {
		if i == tailIdx {
			if digits[i] == 9 {
				plusOne = true
				digits[i] = 0
			} else {
				digits[i] += 1
			}
		} else if plusOne == true {
			if digits[i] == 9 {
				digits[i] = 0
			} else  {
				plusOne = false
				digits[i] += 1
			}
		}
	}


	if plusOne == true {
		result := []int{1}
		for _, n := range digits {
			result = append(result, n)
		}
		return result
	}

	return digits
}