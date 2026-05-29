package stack

import (
	"errors"
)

type node struct {
	value    int
	previous *node
}

func newNode(v int, p *node) *node {
	return &node{value: v, previous: p}
}

type Stack struct {
	head   *node
	length int
}

func NewStack() *Stack {
	return &Stack{head: nil, length: 0}
}

func (s *Stack) Peek() (int, error) {
	if s.length == 0 {
		return -1, errors.New("Stack is empty")
	}

	return s.head.value, nil
}

func (s *Stack) Push(v int) {
	s.length++

	n := newNode(v, nil)
	n.previous = s.head
	s.head = n
}

func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return -1, errors.New("Stack is empty")
	}

	c := s.head
	s.head = s.head.previous
	c.previous = nil

	s.length--

	return c.value, nil
}
