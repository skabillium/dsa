package pathfinding

var directions = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

type Point struct {
	x int
	y int
}

func walk(maze []string, wall byte, curr Point, end Point, path *[]Point, seen [][]bool) bool {
	// Base cases:
	// 1. Is out of bounds
	// 2. Is on wall
	// 3. Is at the end
	// 4. Seen

	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	if maze[curr.y][curr.x] == wall {
		return false
	}

	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}

	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	// Recurse
	for i := 0; i < len(directions); i++ {
		dx, dy := directions[i][0], directions[i][1]
		next := Point{x: curr.x + dx, y: curr.y + dy}
		if walk(maze, wall, next, end, path, seen) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func Solve(maze []string, wall byte, start Point, end Point) []Point {
	path := []Point{}
	seen := make([][]bool, len(maze))
	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[i]))
	}

	walk(maze, wall, start, end, &path, seen)

	return path
}
