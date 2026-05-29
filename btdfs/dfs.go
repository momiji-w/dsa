package btdfs

import "github.com/momiji-w/dsa/utils"

type path struct {
	path []int
}

func (p *path) append(v int) {
	p.path = append(p.path, v)
}

func DFS(node *utils.BinaryNode) []int
