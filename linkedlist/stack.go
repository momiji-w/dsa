package linkedlist

import "errors"

type Stack struct {
	head   *PrevNode
	length int
}

func NewStack() *Stack {
	return &Stack{head: nil, length: 0}
}

func (s *Stack) Peek() (int, error) {
	if s.length == 0 {
		return -1, errors.New("Stack is Empty")
	}

	return s.head.value, nil
}

func (s *Stack) Push(v int) {
	n := NewPrevNode(v, s.head)
	s.head = n
	s.length++
}

func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return -1, errors.New("Stack is Empty")
	}
	ret := s.head.value
	s.head = s.head.previous
	s.length--

	return ret, nil
}
