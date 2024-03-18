package trees

func walkPost(node *BinaryNode, path *[]any) {
	if node == nil {
		return
	}

	walkPost(node.Left, path)
	walkPost(node.Right, path)
	*path = append(*path, node.Value)
}

func BinaryPostOrderTraverse(head *BinaryNode) []any {
	path := &[]any{}
	walkPost(head, path)

	return *path
}
