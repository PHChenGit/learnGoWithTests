package array_and_slice

import (
	"reflect"
	"testing"
)

func TestIntersectionTwo(t *testing.T)  {
	nums1 := []int{1, 2}
	nums2 := []int{1, 1}

	act := IntersectionTwo(nums1, nums2)
	excepted := []int{1}

	if !reflect.DeepEqual(act, excepted) {
		t.Errorf("got %v, but excepted %v", act, excepted)
	}
}
