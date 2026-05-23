package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	test := struct {
		arr      []int
		target   int
		expected int
	}{
		[]int{299, 638, 803, 616, 839, 844, 862, 606, 208, 943},
		844,
		5,
	}

	if res := LinearSearch(test.arr, test.target); res != test.expected {
		t.Fatalf("Expecting %v, got %v", test.expected, res)
	}
}
