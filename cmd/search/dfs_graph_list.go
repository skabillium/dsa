package search

import "skabillium.io/kata-go/cmd/graphs"

func walkGraphList(graph graphs.WeightedAdjacencyList, curr int, needle int, seen []bool, path *[]int) bool {
	if seen[curr] {
		return false
	}
	seen[curr] = true

	*path = append(*path, curr)
	if curr == needle {
		return true
	}

	list := graph[curr]
	for i := 0; i < len(list); i++ {
		if walkGraphList(graph, list[i].To, needle, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func DfsGraphList(graph graphs.WeightedAdjacencyList, source int, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}
	walkGraphList(graph, source, needle, seen, &path)

	if len(path) > 0 {
		return path
	}

	return []int{}
}
