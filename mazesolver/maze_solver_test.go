package mazesolver

import "testing"

func TestMazeSolver(t *testing.T) {
	expected := []Point{
		{x: 1, y: 5},
		{x: 1, y: 4},
		{x: 2, y: 4},
		{x: 3, y: 4},
		{x: 4, y: 4},
		{x: 5, y: 4},
		{x: 6, y: 4},
		{x: 7, y: 4},
		{x: 8, y: 4},
		{x: 9, y: 4},
		{x: 10, y: 4},
		{x: 10, y: 3},
		{x: 10, y: 2},
		{x: 10, y: 1},
		{x: 10, y: 0},
	}
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	path := MazeSolver(maze, 'x', Point{x: 1, y: 5}, Point{x: 10, y: 0})

	if len(path) != len(expected) {
		t.Fatal("Expecting len(path) != len(expected)")
	}

	for i := range path {
		if path[i].x != expected[i].x && path[i].y != expected[i].y {
			t.Fatalf("%d != %d && %d != %d", path[i].x, expected[i].x, path[i].y, expected[i].y)
		}
	}
}
