package trees

func walkPre(node *BinaryNode, path *[]any) {
	if node == nil {
		return
	}

	*path = append(*path, node.value)
	walkPre(node.left, path)
	walkPre(node.right, path)
}

func BinaryPreOrderTraverse(head *BinaryNode) []any {
	path := &[]any{}
	walkPre(head, path)

	return *path
}
