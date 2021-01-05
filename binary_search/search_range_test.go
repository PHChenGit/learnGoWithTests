package binary_search

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T)  {
	t.Run("Case 1", func(t *testing.T) {
		givenNums := []int{5,7,7,8,8,10}
		givenTarget := 8
		excepted := []int{3,4}
		got := searchRange(givenNums, givenTarget)
		shouldBeEqual(got, excepted, t)
	})

	t.Run("Case 2", func(t *testing.T) {
		givenNums := []int{5,7,7,8,8,10}
		givenTarget := 6
		excepted := []int{-1,-1}
		got := searchRange(givenNums, givenTarget)
		shouldBeEqual(got, excepted, t)
	})

	t.Run("Case 3", func(t *testing.T) {
		givenNums := []int{}
		givenTarget := 0
		excepted := []int{-1,-1}
		got := searchRange(givenNums, givenTarget)
		shouldBeEqual(got, excepted, t)
	})

	t.Run("Case 4", func(t *testing.T) {
		givenNums := []int{2,2}
		givenTarget := 2
		excepted := []int{0, 1}
		got := searchRange(givenNums, givenTarget)
		shouldBeEqual(got, excepted, t)
	})

	t.Run("Case 5", func(t *testing.T) {
		givenNums := []int{1,2}
		givenTarget := 2
		excepted := []int{1, 1}
		got := searchRange(givenNums, givenTarget)
		shouldBeEqual(got, excepted, t)
	})

	t.Run("Case 6", func(t *testing.T) {
		givenNums := []int{1,2}
		givenTarget := 1
		excepted := []int{0, 0}
		got := searchRange(givenNums, givenTarget)
		shouldBeEqual(got, excepted, t)
	})
}

func shouldBeEqual(got []int, excepted []int, t *testing.T)  {
	if !reflect.DeepEqual(got, excepted) {
		t.Errorf("got %v, but excepted %v", got, excepted)
	}
}