package array_and_slice

import "testing"

func TestRotate(t *testing.T)  {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	got := Rotate(nums, k)
	excepted := []int{5, 6, 7, 1, 2, 3, 4}

	for i := 0; i < len(excepted); i++ {
		if got[i] != excepted[i] {
			t.Errorf("got %d, but excepte %d", got, excepted)
		}
	}
}
