package binary_search

import "testing"

func TestGuessNumber(t *testing.T)  {
	t.Run("Case 1", func(t *testing.T) {
		given := 10
		excepted := 6
		got := guessNumber(given)

		if got != excepted {
			t.Errorf("got: %d, but excepted: %d", got, excepted)
		}
	})
}
