package two_pointer

import (
	"reflect"
	"testing"
)

func TestSquaresSortedArray(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{-4,-1,0,3,10}
		got := SquaresSortedArray(nums)
		excepted := []int{0,1,9,16,100}

		if !reflect.DeepEqual(got, excepted) {
			t.Errorf("got %d, but excepted %d", got, excepted)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		nums := []int{-4,-1,0,3, 3,10}
		got := SquaresSortedArray(nums)
		excepted := []int{0,1,9,9,16,100}

		if !reflect.DeepEqual(got, excepted) {
			t.Errorf("got %d, but excepted %d", got, excepted)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		nums := []int{-7,-3,2,3,11}
		got := SquaresSortedArray(nums)
		excepted := []int{4,9,9,49,121}

		if !reflect.DeepEqual(got, excepted) {
			t.Errorf("got %d, but excepted %d", got, excepted)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		nums := []int{-5,-3,-2,-1}
		got := SquaresSortedArray(nums)
		excepted := []int{1, 4, 9, 25}

		if !reflect.DeepEqual(got, excepted) {
			t.Errorf("got %d, but excepted %d", got, excepted)
		}
	})
}
