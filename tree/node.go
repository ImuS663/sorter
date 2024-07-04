package tree

import "github.com/ImuS663/sorter"

type node[T sorter.Number] struct {
	Key       T
	LeftNode  *node[T]
	RightNode *node[T]
}

func newNode[T sorter.Number](key T) *node[T] {
	return &node[T]{Key: key}
}

func insert[T sorter.Number](root *node[T], key T) *node[T] {
	if root == nil {
		root = newNode(key)
		return root
	}

	if key < root.Key {
		root.LeftNode = insert(root.LeftNode, key)
	} else {
		root.RightNode = insert(root.RightNode, key)
	}

	return root
}

func inorder[T sorter.Number](root *node[T]) []T {
	var result []T

	if root != nil {
		result = append(result, inorder(root.LeftNode)...)
		result = append(result, root.Key)
		result = append(result, inorder(root.RightNode)...)
	}

	return result
}
