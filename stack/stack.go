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

func (s *Stack) Peek() (int, error)
func (s *Stack) Push(v int)
func (s *Stack) Pop() (int, error)
