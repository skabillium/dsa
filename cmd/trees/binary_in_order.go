package trees

func walkIn(node *BinaryNode, path *[]any) {
	if node == nil {
		return
	}

	walkIn(node.Left, path)
	*path = append(*path, node.Value)
	walkIn(node.Right, path)
}

func BinaryInOrderTraverse(head *BinaryNode) []any {
	path := &[]any{}
	walkIn(head, path)

	return *path
}
