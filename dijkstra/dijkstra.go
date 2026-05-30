package dijkstra

import (
	"math"
	"slices"

	"github.com/momiji-w/dsa/utils"
)

func hasUnvisited(seen []bool, dist []int) bool {
	for i := range seen {
		if !seen[i] && dist[i] != math.MaxInt {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dist []int) int {
	lowest := math.MaxInt
	lowestIdx := -1
	for i := range seen {
		if seen[i] {
			continue
		}

		if dist[i] < lowest {
			lowest = dist[i]
			lowestIdx = i
		}
	}
	return lowestIdx
}

func Dijkstra(graph utils.Graph, source, target int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}
	dist := make([]int, len(graph))
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[source] = 0

	for hasUnvisited(seen, dist) {
		n := getLowestUnvisited(seen, dist)
		seen[n] = true
		adjs := graph[n]

		for i := range adjs {
			e := adjs[i]
			curW := dist[n] + e.Weight
			if dist[e.To] > curW {
				dist[e.To] = curW
				prev[e.To] = n
			}
		}
	}

	path := make([]int, 0)
	cur := target
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
