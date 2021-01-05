package binary_search

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

const (
	isLower int = -1
	isHigher int = 1
	isEqual int = 0
	ans int = 6
)

func guessNumber(n int) int {
	left := 1
	right := n
	var res int
	var mid int

	for left < right {
		mid = left + (right - left) / 2
		res =  guess(mid)

		if res == isEqual {
			break
		} else if res == isHigher {
			right = mid - 1
		} else if res == isLower {
			left = mid + 1
		}
	}

	return mid
}

func guess(n int) int {
	if n < ans {
		return isLower
	} else if n > ans {
		return isHigher
	}

	return isEqual
}