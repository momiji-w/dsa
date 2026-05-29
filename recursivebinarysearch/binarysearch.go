package recursivebinarysearch

func binarySearch(arr []int, target int, lo, hi int) int {
	if lo > hi {
		return -1
	}

	m := lo + (hi-lo)/2

	if arr[m] == target {
		return m
	}

	if arr[m] > target {
		return binarySearch(arr, target, lo, m-1)
	} else {
		return binarySearch(arr, target, m+1, hi)
	}
}

func BinarySearch(arr []int, target int) int {
	return binarySearch(arr, target, 0, len(arr)-1)
}
