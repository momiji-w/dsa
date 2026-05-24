package btbfs

import (
	"github.com/momiji-w/dsa/utils"
	"slices"
	"testing"
)

func TestBFS(t *testing.T) {
	expected := []int{10, 11, 12, 13, 14, 15, 16}

	path := BFS(utils.Root)

	if !slices.Equal(path, expected) {
		t.Fatalf("%v != %v\n", path, expected)
	}
}
