package btdfs

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

func TestDFS(t *testing.T) {
	expected := []int{13, 11, 14, 10, 15, 12, 16}

	path := DFS(utils.Root)

	if !slices.Equal(path, expected) {
		t.Fatalf("%v != %v\n", path, expected)
	}
}

func TestCompare(t *testing.T) {
	if !Compare(root, root) {
		t.Fatalf("Tree should be identical")
	}

	if Compare(root, root2) {
		t.Fatalf("Tree not should be identical")
	}
}
