package search

import (
	"skabillium.io/kata-go/cmd/structures"
	"skabillium.io/kata-go/cmd/trees"
)

func BinaryBfs(head trees.BinaryNode, needle any) bool {
	q := structures.NewQueue()
	q.Enqueue(head)

	for q.Length != 0 {
		node, _ := q.Deque().(trees.BinaryNode)
		if node.Value == needle {
			return true
		}

		if node.Left != nil {
			q.Enqueue(*node.Left)
		}
		if node.Right != nil {
			q.Enqueue(*node.Right)
		}
	}

	return false
}
