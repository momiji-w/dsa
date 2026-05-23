package linkedlist

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(100)
	q.Enqueue(101)
	p, err := q.Peek()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	if p != 100 {
		t.Fatalf("Expecting 100 from Peek, got %d", p)
	}

	d, err := q.Dequeue()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	if d != 100 {
		t.Fatalf("Expecting 100 from Dequeue, got %d", d)
	}

	p, err = q.Peek()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	if p != 101 {
		t.Fatalf("Expecting 101 from Peek, got %d", p)
	}

	d, err = q.Dequeue()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	if d != 101 {
		t.Fatalf("Expecting 101 from Dequeue, got %d", d)
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Fatalf("Expecting queue to be empty")
	}
}
