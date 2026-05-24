package graphdfs

import (
	"github.com/momiji-w/dsa/utils"
	"slices"
	"testing"
)

func TestDFS(t *testing.T) {
	res := DFS(utils.G, 0, 6)
	expected := []int{0, 1, 2, 3, 4, 5, 6}
	if !slices.Equal(res, expected) {
		t.Fatalf("expected %v, got %v", expected, res)
	}
}
