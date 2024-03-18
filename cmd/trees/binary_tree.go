package trees

type BinaryNode struct {
	Value any
	Left  *BinaryNode
	Right *BinaryNode
}

var ExampleBinaryTree = &BinaryNode{
	Value: 20,
	Right: &BinaryNode{
		Value: 50,
		Right: &BinaryNode{
			Value: 100,
		},
		Left: &BinaryNode{
			Value: 30,
			Right: &BinaryNode{
				Value: 45,
			},
			Left: &BinaryNode{
				Value: 29,
			},
		},
	},
	Left: &BinaryNode{
		Value: 10,
		Right: &BinaryNode{
			Value: 15,
		},
		Left: &BinaryNode{
			Value: 5,
			Right: &BinaryNode{
				Value: 7,
			},
		},
	},
}

var ExampleBinaryTree2 = &BinaryNode{
	Value: 20,
	Right: &BinaryNode{
		Value: 50,
		Left: &BinaryNode{
			Value: 30,
			Right: &BinaryNode{
				Value: 45,
				Right: &BinaryNode{
					Value: 49,
				},
			},
			Left: &BinaryNode{
				Value: 29,
				Left: &BinaryNode{
					Value: 21,
				},
			},
		},
	},
	Left: &BinaryNode{
		Value: 10,
		Right: &BinaryNode{
			Value: 15,
		},
		Left: &BinaryNode{
			Value: 5,
			Right: &BinaryNode{
				Value: 7,
			},
		},
	},
}
