package array_and_slice

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{1,2,3}
		got := PlusOne(nums)
		excepted :=  []int{1,2,4}

		if !reflect.DeepEqual(got, excepted) {
			t.Errorf("got %v, but excepted %v", got, excepted)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		nums := []int{9,9}
		got := PlusOne(nums)
		excepted :=  []int{1,0,0}

		if !reflect.DeepEqual(got, excepted) {
			t.Errorf("got %v, but excepted %v", got, excepted)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		nums := []int{1,9,4}
		got := PlusOne(nums)
		excepted :=  []int{1,9,5}

		if !reflect.DeepEqual(got, excepted) {
			t.Errorf("got %v, but excepted %v", got, excepted)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		nums := []int{4,3,2,1}
		got := PlusOne(nums)
		excepted :=  []int{4,3,2,2}

		if !reflect.DeepEqual(got, excepted) {
			t.Errorf("got %v, but excepted %v", got, excepted)
		}
	})
}
