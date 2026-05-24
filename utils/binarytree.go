package utils

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

func NewBinaryNode(v int, l, r *BinaryNode) *BinaryNode {
	return &BinaryNode{v, l, r}
}

var Root = NewBinaryNode(10,
	NewBinaryNode(11,
		NewBinaryNode(13, nil, nil),
		NewBinaryNode(14, nil, nil)),
	NewBinaryNode(12,
		NewBinaryNode(15, nil, nil),
		NewBinaryNode(16, nil, nil)),
)

var Root2 = NewBinaryNode(10,
	NewBinaryNode(11,
		NewBinaryNode(13, nil, nil),
		NewBinaryNode(14, nil, nil)),
	NewBinaryNode(12, NewBinaryNode(15, nil, nil), nil),
)
