package dijkstra

import (
	"github.com/momiji-w/dsa/utils"
	"slices"
	"testing"
)

func TestDijkstra(t *testing.T) {
	res := Dijkstra(utils.G, 0, 6)
	expected := []int{0, 1, 4, 5, 6}
	if !slices.Equal(res, expected) {
		t.Fatalf("expected %v, got %v", expected, res)
	}
}
