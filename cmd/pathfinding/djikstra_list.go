package pathfinding

import (
	"math"
	"slices"

	"skabillium.io/kata-go/cmd/graphs"
)

func hasUnvisited(seen []bool, dists []int) bool {
	for i := 0; i < len(seen); i++ {
		if !seen[i] && dists[i] < math.MaxInt {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	min := math.MaxInt
	for i := 0; i < len(seen); i++ {
		if !seen[i] && dists[i] < min {
			idx = i
			min = dists[i]
		}
	}

	return idx
}

func DjikstraList(graph graphs.WeightedAdjacencyList, source int, sink int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	dists := make([]int, len(graph))

	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}

	for i := 0; i < len(dists); i++ {
		dists[i] = math.MaxInt
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge.To] {
				continue
			}

			dist := dists[edge.To] + edge.Weight
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	path := []int{}
	curr := sink
	for prev[curr] != -1 {
		path = append(path, curr)
		curr = prev[curr]
	}

	path = append(path, source)
	slices.Reverse(path)

	return path
}
