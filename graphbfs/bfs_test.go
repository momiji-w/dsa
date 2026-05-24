package graphbfs

import (
	"testing"
	"slices"
	"github.com/momiji-w/dsa/utils"
)

func TestBFS(t *testing.T) {
        res := BFS(utils.G, 0, 6)
        expected := []int{0, 2, 3, 6}
        if !slices.Equal(res, expected) {
                t.Fatalf("expected %v, got %v", expected, res)
        }
}
