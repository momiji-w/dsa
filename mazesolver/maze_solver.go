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

func walk(maze []string, wall byte, cur, end Point, seen [][]bool, path *Path) bool {
	if cur.y >= len(maze) || cur.x >= len(maze[0]) {
		return false
	}

	if maze[cur.y][cur.x] == wall {
		return false
	}

	if cur.x == end.x && cur.y == end.y {
		path.Append(end)
		return true
	}

	if seen[cur.y][cur.x] {
		return false
	}

	seen[cur.y][cur.x] = true
	path.Append(cur)

	for _, d := range directions {
		if walk(maze, wall, Point{x: cur.x + d.x, y: cur.y + d.y}, end, seen, path) {
			return true
		}
	}

	path.Pop()

	return false
}

func MazeSolver(maze []string, wall byte, start, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[i]))
	}
	path := &Path{path: make([]Point, 0)}
	walk(maze, wall, start, end, seen, path)
	return path.path
}
