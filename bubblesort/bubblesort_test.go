package bubblesort

import (
	"github.com/momiji-w/dsa/utils"
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, v := range utils.SortingTest {
		k := v.Arr
		BubbleSort(k)
		if !slices.Equal(k, v.Expected) {
			t.Errorf("Topic: %s, expecting %v, got %v", v.Topic, v.Expected, k)
		}
	}
}
