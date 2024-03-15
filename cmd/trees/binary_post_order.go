package trees

func walkPost(node *BinaryNode, path *[]any) {
	if node == nil {
		return
	}

	walkPost(node.left, path)
	walkPost(node.right, path)
	*path = append(*path, node.value)
}

func BinaryPostOrderTraverse(head *BinaryNode) []any {
	path := &[]any{}
	walkPost(head, path)

	return *path
}
