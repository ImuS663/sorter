package tree

import "github.com/ImuS663/sorter"

type Tree[T sorter.Number] struct{}

func (t Tree[T]) Sort(arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	var root *node[T]

	for _, item := range arr {
		root = insert(root, item)
	}

	return inorder(root)
}
