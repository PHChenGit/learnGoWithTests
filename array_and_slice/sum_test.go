package array_and_slice

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	excepted := 15

	if got != excepted {
		t.Errorf("got %d, excepted %d", got, excepted)
	}
}
