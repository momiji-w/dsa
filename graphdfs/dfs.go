package graphdfs

import (
	"github.com/momiji-w/dsa/utils"
)

type Path struct {
	path []int
}

func (p *Path) Append(node int) {
	p.path = append(p.path, node)
}

func (p *Path) Pop() {
	p.path = p.path[:len(p.path)-1]
}

func walk(graph utils.Graph, cur, target int, path *Path, seen []bool) bool {
	if seen[cur] {
		return false
	}

	if cur == target {
		path.Append(cur)
		return true
	}

	path.Append(cur)
	seen[cur] = true

	for _, e := range graph[cur] {
		if walk(graph, e.To, target, path, seen) {
			return true
		}
	}

	path.Pop()
	return false
}

func DFS(graph utils.Graph, source, target int) []int {
	seen := make([]bool, len(graph))
	path := &Path{path: make([]int, 0)}

	walk(graph, source, target, path, seen)

	return path.path
}
