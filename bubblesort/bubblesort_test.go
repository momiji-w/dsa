package bubblesort

import (
	"github.com/momiji-w/dsa/utils"
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, v := range utils.TestTopics {
		k := v.Arr
		BubbleSort(k)
		if !slices.Equal(k, v.Expected) {
		}
	}
}
