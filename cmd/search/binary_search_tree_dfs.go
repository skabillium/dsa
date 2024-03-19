package search

import (
	"skabillium.io/kata-go/cmd/trees"
)

func bstDfs(curr *trees.BinaryNode, needle any) bool {
	if curr == nil {
		return false
	}

	if curr.Value == needle {
		return true
	}

	nValue, _ := needle.(int)
	cValue, _ := curr.Value.(int)

	if nValue < cValue {
		return bstDfs(curr.Left, needle)
	}

	return bstDfs(curr.Right, needle)
}

func BinarySearchTreeDfs(head *trees.BinaryNode, needle int) bool {
	return bstDfs(head, needle)
}
