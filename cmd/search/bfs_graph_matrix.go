package search

import (
	"slices"

	"skabillium.io/kata-go/cmd/graphs"
	"skabillium.io/kata-go/cmd/structures"
)

func BfsGraphMatrix(graph graphs.WeightedAdjacencyMatrix, source int, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))

	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}

	seen[source] = true
	q := structures.NewQueue()
	q.Enqueue(source)

	for {
		curr := q.Deque().(int)
		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q.Enqueue(i)
		}

		if q.Length == 0 {
			break
		}
	}

	// Build it backwards
	curr := needle
	out := []int{}
	for curr != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) != 0 {
		slices.Reverse(out)
		return out
	}

	return []int{}
}
