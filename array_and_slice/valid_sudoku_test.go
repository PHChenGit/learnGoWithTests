package array_and_slice

import "testing"

func TestValidSudoku(t *testing.T)  {
	t.Run("Test valid sudo", func(t *testing.T) {
		//board := [9][9]byte{
		//	{'5','3','.','.','7','.','.','.','.'},
		//	{'6','.','.','1','9','5','.','.','.'},
		//	{'.','9','8','.','.','.','.','6','.'},
		//	{'8','.','.','.','6','.','.','.','3'},
		//	{'4','.','.','8','.','3','.','.','1'},
		//	{'7','.','.','.','2','.','.','.','6'},
		//	{'.','6','.','.','.','.','2','8','.'},
		//	{'.','.','.','4','1','9','.','.','5'},
		//	{'.','.','.','.','8','.','.','7','9'},
		//}
		board := make([][]byte, 0, 0)
		row1 :=  []byte{'5','3','.','.','7','.','.','.','.'}
		board = append(board, row1)

		got := ValidSudoku(board)
		excepted := true

		if got != excepted {
			t.Errorf("got %t, excepted %t", got, excepted)
		}
	})
}
