package quicksort

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if pivot > arr[j] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[hi] = arr[i]
	arr[i] = pivot

	return i
}

func quicksort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := partition(arr, lo, hi)

	quicksort(arr, lo, p-1)
	quicksort(arr, p+1, hi)
}

func QuickSort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}
