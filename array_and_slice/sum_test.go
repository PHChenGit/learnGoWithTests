package array_and_slice

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		excepted := 15

		if got != excepted {
			t.Errorf("got %d, excepted %d", got, excepted)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		excepted := 6

		if got != excepted {
			t.Errorf("got %d, excepted %d", got, excepted)
		}
	})
}
