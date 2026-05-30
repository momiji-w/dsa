package btbfs

import (
	"github.com/momiji-w/dsa/utils"
	"container/list"
)

func BFS(node *utils.BinaryNode) []int {
	q := list.New()
	q.PushBack(node)
	path := make([]int, 0)

	for q.Len() > 0 {
		n := q.Remove(q.Front()).(*utils.BinaryNode)
		path = append(path, n.Value)

		if n.Left != nil {
			q.PushBack(n.Left)
		}
		if n.Right != nil {
			q.PushBack(n.Right)
		}
	}

	return path
}
