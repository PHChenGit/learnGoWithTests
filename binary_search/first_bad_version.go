package binary_search

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

const (
	firstBadVersionAns int = 4
)

func firstBadVersion(n int) int {
	left := 1
	right := n
	var mid int

	for left < right {
		mid = left + (right - left) / 2

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func isBadVersion(num int) bool {
	if num >= firstBadVersionAns {
		return true
	}

	return false
}
