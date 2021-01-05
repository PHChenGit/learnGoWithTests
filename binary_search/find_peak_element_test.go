package binary_search

import "testing"

func TestFindPeakElement(t *testing.T)  {
	t.Run("Case 1", func(t *testing.T) {
		given := []int{1,2,3,1}
		excepted := 2
		got := findPeakElement(given)

		shouldPass(got, excepted, t)
	})

	t.Run("Case 2", func(t *testing.T) {
		given := []int{1,2,1,3,5,6,4}
		excepted := 5
		got := findPeakElement(given)

		shouldPass(got, excepted, t)
	})

	t.Run("Case 3", func(t *testing.T) {
		given := []int{6,5,4,3,2,3,2}
		excepted := 0
		got := findPeakElement(given)

		shouldPass(got, excepted, t)
	})

	t.Run("Case 4", func(t *testing.T) {
		given := []int{2,1}
		excepted := 0
		got := findPeakElement(given)

		shouldPass(got, excepted, t)
	})

	t.Run("Case 5", func(t *testing.T) {
		given := []int{1,2}
		excepted := 1
		got := findPeakElement(given)

		shouldPass(got, excepted, t)
	})

	t.Run("Case 6", func(t *testing.T) {
		given := []int{1,2,3,4,5,6}
		excepted := 5
		got := findPeakElement(given)

		shouldPass(got, excepted, t)
	})
}

func shouldPass(got int, excepted int, t *testing.T) {
	if got != excepted {
		t.Errorf("got %d, but excepted %d", got, excepted)
	}
}