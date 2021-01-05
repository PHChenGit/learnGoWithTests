package array_and_slice

import "testing"

func TestBuyAndSell(t *testing.T)  {
	//prices := []int{7,1,5,3,6,4}

	shouldPass := func(prices []int, excepted int) {
		got := BuyAndSellTwo(prices)

		if got != excepted {
			t.Errorf("got %d, but excepted %d", got, excepted)
		}
	}

	t.Run("Case 1", func(t *testing.T) {
		prices := []int{7,1,5,3,6,4}
		shouldPass(prices, 7)
	})

	t.Run("Case 2", func(t *testing.T) {
		prices := []int{3,3}
		shouldPass(prices, 0)
	})
}