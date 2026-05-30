package graphbfs

import (
	"container/list"
	"slices"

	"github.com/momiji-w/dsa/utils"
)

func BFS(graph utils.Graph, source, target int) []int {
	q := list.New()
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	q.PushBack(source)
	seen[source] = true
	for q.Len() > 0 {
		cur := q.Remove(q.Front()).(int)
		for _, e := range graph[cur] {
			if seen[e.To] {
				continue
			}

			seen[e.To] = true
			prev[e.To] = cur

			q.PushBack(e.To)
		}
	}

	cur := target
	path := make([]int, 0)
	if prev[cur] == -1 {
		return path
	}

	for prev[cur] != -1 {
		path = append(path, cur)
		cur = prev[cur]
	}
	path = append(path, source)
	slices.Reverse(path)

	return path
}
