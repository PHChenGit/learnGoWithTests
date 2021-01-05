package array_and_slice

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicationNumbers(t *testing.T)  {
	numbers := []int{0,0,1,1,1,2,2,3,3,4}
	got := RemoveDuplicationNumbers(numbers)
	//excepted := []int{0, 1, 2, 3, 4}
	excepted := 5

	if !reflect.DeepEqual(got, excepted) {
		t.Errorf("got %v, excepted %v", got, excepted)
	}
}
