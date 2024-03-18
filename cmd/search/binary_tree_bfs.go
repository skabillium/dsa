package search

import (
	"skabillium.io/kata-go/cmd/structures"
	"skabillium.io/kata-go/cmd/trees"
)

// func prepend(arr []trees.BinaryNode, node trees.BinaryNode) []trees.BinaryNode {
// 	newArr := make([]trees.BinaryNode, len(arr)+1)
// 	copy(newArr[1:], arr)
// 	newArr[0] = node

// 	return newArr
// }

// func pop(arr []trees.BinaryNode) trees.BinaryNode {
// 	node := arr[len(arr)-1]
// 	arr = arr[:len(arr)-1]

// 	return node
// }

// TODO: Figure out how to use the queue created before
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
