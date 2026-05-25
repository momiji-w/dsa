package recursivebinarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	test := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{5, 21, 38, 48, 58, 492}, 492, 5},
		{[]int{5, 21, 38, 48, 58, 492}, 5, 0},
		{[]int{5, 21, 38, 48, 58, 492}, 38, 2},

		{[]int{5, 21, 38, 48, 58, 492}, 8, -1},

		{[]int{}, 8, -1},
	}

	for _, v := range test {
		if res := BinarySearch(v.arr, v.target); res != v.expected {
			t.Fatalf("Binary Search: expected %d, got %d", v.expected, res)
		}
	}
}
