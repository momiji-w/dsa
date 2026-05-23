package linkedlist

import (
	"errors"
	"fmt"
)

type Queue struct {
	head   *NextNode
	tail   *NextNode
	length int
}

func NewQueue() *Queue {
	return &Queue{head: nil, tail: nil, length: 0}
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Peek() (int, error) {
	if q.length == 0 {
		return -1, errors.New("Queue is Empty")
	}

	return q.head.value, nil
}

func (q *Queue) Enqueue(v int) {
	n := NewNextNode(v)
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
	fmt.Println(q.length)
	if q.length == 0 {
		return -1, errors.New("Queue is Empty")
	}
	q.length--
	fmt.Println(q.length)
	ret := q.head.value
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return ret, nil
}
