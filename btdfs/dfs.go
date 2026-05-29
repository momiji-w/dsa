package btdfs

import "github.com/momiji-w/dsa/utils"

type path struct {
	path []int
}

func (p *path) append(v int) {
	p.path = append(p.path, v)
}

func walk(node *utils.BinaryNode, path *path) {
	if node == nil {
		return
	}

	walk(node.Left, path)
	path.append(node.Value)
	walk(node.Right, path)
}

func DFS(node *utils.BinaryNode) []int {
	path := &path{path: make([]int, 0)}
	walk(node, path)
	return path.path
}
