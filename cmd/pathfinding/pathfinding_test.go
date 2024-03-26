package pathfinding_test

import (
	"fmt"
	"reflect"
	"testing"

	"skabillium.io/kata-go/cmd/graphs"
	"skabillium.io/kata-go/cmd/pathfinding"
)

func TestDjikstraList(t *testing.T) {
	expected := []int{0, 1, 4, 5, 6}
	res := pathfinding.DjikstraList(graphs.List1, 0, 6)
	if !reflect.DeepEqual(res, expected) {
		t.Error("Expected DjikstraList(graphs.List1, 0, 6) to return", expected, "got", res)
	}
}

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	expected := []pathfinding.Point{
		{X: 10, Y: 0},
		{X: 10, Y: 1},
		{X: 10, Y: 2},
		{X: 10, Y: 3},
		{X: 10, Y: 4},
		{X: 9, Y: 4},
		{X: 8, Y: 4},
		{X: 7, Y: 4},
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 2, Y: 4},
		{X: 1, Y: 4},
		{X: 1, Y: 5},
	}

	result := pathfinding.SolveMaze(maze, 'x', pathfinding.Point{X: 10, Y: 0}, pathfinding.Point{X: 1, Y: 5})
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected other path")
		fmt.Println("Received:", result)
		fmt.Println("Expected:", expected)
	}
}
