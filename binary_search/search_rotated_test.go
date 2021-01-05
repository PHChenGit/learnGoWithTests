package binary_search

import "testing"

func TestSearchRotatedArray(t *testing.T)  {
	t.Run("Case 1", func(t *testing.T) {
		givenArr := []int{4,5,6,7,0,1,2}
		givenTarget := 5
		excepted := 1
		got := SearchRotatedArray(givenArr, givenTarget)

		if got != excepted {
			t.Errorf("got: %d, but excepted: %d", got, excepted)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		givenArr := []int{4,5,6,7,0,1,2}
		givenTarget := 0
		excepted := 4
		got := SearchRotatedArray(givenArr, givenTarget)

		if got != excepted {
			t.Errorf("got: %d, but excepted: %d", got, excepted)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		givenArr := []int{1}
		givenTarget := 0
		excepted := -1
		got := SearchRotatedArray(givenArr, givenTarget)

		if got != excepted {
			t.Errorf("got: %d, but excepted: %d", got, excepted)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		givenArr := []int{5,1,3}
		givenTarget := 5
		excepted := 0
		got := SearchRotatedArray(givenArr, givenTarget)

		if got != excepted {
			t.Errorf("got: %d, but excepted: %d", got, excepted)
		}
	})

	t.Run("Case 5", func(t *testing.T) {
		givenArr := []int{1,2}
		givenTarget := 2
		excepted := 1
		got := SearchRotatedArray(givenArr, givenTarget)

		if got != excepted {
			t.Errorf("got: %d, but excepted: %d", got, excepted)
		}
	})

	t.Run("Case 6", func(t *testing.T) {
		givenArr := []int{1,2}
		givenTarget := 1
		excepted := 0
		got := SearchRotatedArray(givenArr, givenTarget)

		if got != excepted {
			t.Errorf("got: %d, but excepted: %d", got, excepted)
		}
	})
}