package linkedlist

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(100)
	p, err := s.Peek()

	if err != nil{
		t.Fatalf("%s", err.Error())
	}
	if p != 100 {
		t.Fatalf("Expecting 100, got %d", p)
	}

	k, err := s.Pop()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if k != 100 {
		t.Fatalf("Expecting Pop 100, got %d", k)
	}

	_, err = s.Pop()
	if err == nil {
		t.Fatalf("Expecting stack to be empty")
	}
}
