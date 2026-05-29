package linkedlist

import (
	"errors"
)

type node struct {
	value int
	next  *node
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

func (q *Queue) Peek() (int, error) {
	if q.length == 0 {
		return -1, errors.New("Queue is empty")
	}

	return q.head.value, nil
}

func (q *Queue) Enqueue(v int) {
	n := newNode(v)
	q.length++
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *Queue) Dequeue() (int, error) {
	if q.length == 0 {
		return -1, errors.New("Queue is empty")
	}

	q.length--
	c := q.head

	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		c.next = nil
		return c.value, nil
	}

	q.head = q.head.next
	c.next = nil

	return c.value, nil
}
