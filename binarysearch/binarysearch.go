package binarysearch

func BinarySearch(arr []int, target int) int {
	hi := len(arr) - 1
	lo := 0

	for hi >= lo {
		m := lo + ((hi - lo) / 2)

		if arr[m] == target {
			return m
		}

		if arr[m] > target {
			hi = m - 1
		} else {
			lo = m + 1
		}
	}

	return -1
}
