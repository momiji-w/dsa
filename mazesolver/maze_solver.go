package mazesolver

type Point struct {
	x int
	y int
}

var directions = [4]Point{
	{x: 0, y: 1},
	{x: 1, y: 0},
	{x: 0, y: -1},
	{x: -1, y: 0},
}

type Path struct {
	path []Point
}

func (p *Path) Append(point Point) {
	p.path = append(p.path, point)
}

func (p *Path) Pop() {
	p.path = p.path[:len(p.path)-1]
}

func MazeSolver(maze []string, wall byte, start, end Point) []Point
