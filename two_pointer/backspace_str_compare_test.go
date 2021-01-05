package two_pointer

import "testing"

func TestBackspaceStrCompare(t *testing.T)  {
	t.Run("Case 1", func(t *testing.T) {
		S := "ab#c"
		T := "ad#c"

		got := BackspaceStrCompare(S, T)
		excepted := true

		if got != excepted {
			t.Errorf("got %t, excepted %t", got, excepted)
		}
	})
}
