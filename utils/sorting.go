package utils

type sortTest struct {
	Topic    string
	Arr      []int
	Expected []int
}

var TestTopics = []sortTest{
	{
		Topic:    "Regular unsorted array",
		Arr:      []int{5, 2, 9, 1, 5, 6},
		Expected: []int{1, 2, 5, 5, 6, 9},
	},
	{
		Topic:    "Already sorted array",
		Arr:      []int{1, 2, 3, 4, 5},
		Expected: []int{1, 2, 3, 4, 5},
	},
	{
		Topic:    "Reverse sorted array",
		Arr:      []int{10, 8, 6, 4, 2},
		Expected: []int{2, 4, 6, 8, 10},
	},
	{
		Topic:    "Array with negative numbers",
		Arr:      []int{3, -1, 0, -5, 2},
		Expected: []int{-5, -1, 0, 2, 3},
	},
	{
		Topic:    "Alternating high and low numbers",
		Arr:      []int{100, 1, 99, 2, 98, 3},
		Expected: []int{1, 2, 3, 98, 99, 100},
	},
	{
		Topic:    "Single element array",
		Arr:      []int{42},
		Expected: []int{42},
	},
	{
		Topic:    "Empty array",
		Arr:      []int{},
		Expected: []int{},
	},
	{
		Topic:    "Nil array (uninitialized slice)",
		Arr:      nil,
		Expected: nil,
	},
}
