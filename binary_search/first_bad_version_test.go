package binary_search

import "testing"

func TestFirstBadVersion(t *testing.T)  {
	given := 10
	excepted := 4

	got := firstBadVersion(given)

	if got != excepted {
		t.Errorf("got: %d, but excepted: %d", got, excepted)
	}
}
