package array_and_slice

func IntersectionTwo(nums1 []int, nums2 []int) []int {
	tables := make(map[int]int)
	for _, n := range nums1 {
		if v, ok := tables[n]; ok {
			tables[n] = v + 1
		} else {
			tables[n] = 1
		}
	}

	var result []int
	for _, n := range nums2 {
		if v, ok := tables[n]; ok {
			if v > 0 {
				tables[n] = v - 1
				result = append(result, n)
			}
		}
	}

	return result
}
