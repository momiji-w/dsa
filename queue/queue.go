package linkedlist

import (
	"errors"
)

type node struct {
	value int
	next *node
}

func newNode(v int) *node {
	return &node{value: v, next: nil}
}

type Queue struct {
	head   *node
	tail   *node
	length int
}

func NewQueue() *Queue {
	return &Queue{head: nil, tail: nil, length: 0}
}

func (q *Queue) Peek() (int, error)
func (q *Queue) Enqueue(v int)
func (q *Queue) Dequeue() (int, error)
