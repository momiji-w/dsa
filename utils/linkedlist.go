package utils

type PrevNode struct {
	value    int
	previous *PrevNode
}

func NewPrevNode(v int, p *PrevNode) *PrevNode {
	return &PrevNode{value: v, previous: p}
}

type NextNode struct {
	value int
	next *NextNode
}

func NewNextNode(v int) *NextNode {
	return &NextNode{value: v, next: nil}
}
