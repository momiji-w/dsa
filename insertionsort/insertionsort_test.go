package insertionsort

import (
	"github.com/momiji-w/dsa/utils"
	"slices"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, v := range utils.TestTopics {
		k := v.Arr
		InsertionSort(k)
		if !slices.Equal(k, v.Expected) {
		}
	}
}
