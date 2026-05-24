package heap

import (
	"slices"
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := &Heap{nodes: []int{50, 100, 70}}

	heap.Insert(3)
	insertExpected := []int{3, 50, 70, 100}
	if !slices.Equal(heap.nodes, insertExpected) {
		t.Fatalf("Insert 3: %v != %v", heap.nodes, insertExpected)
	}

	d := heap.Delete()
	if d != 3 {
		t.Fatalf("Delete: Expecting 3, got %v", d)
	}

	deleteExpected := []int{50, 100, 70}
	if !slices.Equal(heap.nodes, deleteExpected) {
		t.Fatalf("Insert 3: %v != %v", heap.nodes, deleteExpected)
	}

	for range heap.Len() {
		heap.Delete()
	}

	if heap.Len() != 0 {
		t.Fatal("Heap should be empty")
	}

	d = heap.Delete()
	if d != -1 {
		t.Fatal("Heap should be empty")
	}
}
