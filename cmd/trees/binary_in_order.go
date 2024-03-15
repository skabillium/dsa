package trees

func walkIn(node *BinaryNode, path *[]any) {
	if node == nil {
		return
	}

	walkIn(node.left, path)
	*path = append(*path, node.value)
	walkIn(node.right, path)
}

func BinaryInOrderTraverse(head *BinaryNode) []any {
	path := &[]any{}
	walkIn(head, path)

	return *path
}
