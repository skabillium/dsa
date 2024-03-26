package pathfinding

var directions = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func walk(maze []string, wall byte, curr Point, end Point, path *[]Point, seen [][]bool) bool {
	// Base cases:
	// 1. Is out of bounds
	// 2. Is on wall
	// 3. Is at the end
	// 4. Seen

	if curr.X < 0 || curr.X >= len(maze[0]) || curr.Y < 0 || curr.Y >= len(maze) {
		return false
	}

	if maze[curr.Y][curr.X] == wall {
		return false
	}

	if curr.X == end.X && curr.Y == end.Y {
		*path = append(*path, end)
		return true
	}

	if seen[curr.Y][curr.X] {
		return false
	}

	seen[curr.Y][curr.X] = true
	*path = append(*path, curr)

	// Recurse
	for i := 0; i < len(directions); i++ {
		dx, dy := directions[i][0], directions[i][1]
		next := Point{X: curr.X + dx, Y: curr.Y + dy}
		if walk(maze, wall, next, end, path, seen) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func SolveMaze(maze []string, wall byte, start Point, end Point) []Point {
	path := []Point{}
	seen := make([][]bool, len(maze))
	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[i]))
	}

	walk(maze, wall, start, end, &path, seen)

	return path
}
