package trees

func walkPre(node *BinaryNode, path *[]any) {
	if node == nil {
		return
	}

	*path = append(*path, node.Value)
	walkPre(node.Left, path)
	walkPre(node.Right, path)
}

func BinaryPreOrderTraverse(head *BinaryNode) []any {
	path := &[]any{}
	walkPre(head, path)

	return *path
}
