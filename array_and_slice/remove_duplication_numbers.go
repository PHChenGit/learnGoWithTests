package array_and_slice

func RemoveDuplicationNumbers(numbers []int) int {
	i := 0
	for j := 1; j < len(numbers); j++ {
		if numbers[i] != numbers[j] {
			i += 1
			numbers[i] = numbers[j]
		}
	}
	return i + 1
}
