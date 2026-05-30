package heapsort

func parent(idx int) int     { return (idx - 1) / 2 }
func leftChild(idx int) int  { return (idx * 2) + 1 }
func rightChild(idx int) int { return (idx * 2) + 2 }

func heapifyDown(arr []int, idx, end int) {
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if idx > end || lIdx > end {
		return
	}

	maxIdx := lIdx
	if rIdx <= end && arr[lIdx] < arr[rIdx] {
		maxIdx = rIdx
	}

	if maxIdx <= end && arr[maxIdx] > arr[idx] {
		arr[maxIdx], arr[idx] = arr[idx], arr[maxIdx]
		heapifyDown(arr, maxIdx, end)
	}
}

func Heapsort(arr []int) {
	for p := parent(len(arr) - 1); p >= 0; p-- {
		heapifyDown(arr, p, len(arr)-1)
	}
	end := len(arr) - 1
	for end > 0 {
		arr[end], arr[0] = arr[0], arr[end]
		end--
		heapifyDown(arr, 0, end)
	}
}
