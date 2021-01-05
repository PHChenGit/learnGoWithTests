package two_pointer

func SquaresSortedArray(nums []int) []int {
	n := []int{}

	for _, num := range nums{
		n = append(n, num*num)
	}

	i := 0
	j := len(n) - 1

	for i < j {
		if n[i] <= n[j] {
			j -= 1
		} else {
			n[i], n[j] = n[j], n[i]
		}
	}

	return n
}