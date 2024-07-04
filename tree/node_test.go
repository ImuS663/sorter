package tree

import (
	"reflect"
	"testing"
)

func Test_newNode(t *testing.T) {
	result := newNode(10)

	if result == nil {
		t.Errorf("expected <no nil>, got %v", result)
	}
}

func Test_insert(t *testing.T) {
	var root *node[int]

	root = insert(root, 4)
	root = insert(root, 2)
	root = insert(root, -2)
	root = insert(root, 6)
	root = insert(root, 3)

	if root.Key != 4 {
		t.Errorf("expected %v, got %v", 4, root.Key)
	}

	if root.LeftNode.Key != 2 {
		t.Errorf("expected %v, got %v", 2, root.LeftNode.Key)
	}

	if root.LeftNode.LeftNode.Key != -2 {
		t.Errorf("expected %v, got %v", -2, root.LeftNode.LeftNode.Key)
	}

	if root.LeftNode.RightNode.Key != 3 {
		t.Errorf("expected %v, got %v", 3, root.LeftNode.RightNode.Key)
	}

	if root.RightNode.Key != 6 {
		t.Errorf("expected %v, got %v", 6, root.RightNode.Key)
	}
}

func Test_inorder(t *testing.T) {
	root := &node[int]{
		Key: 4,
		LeftNode: &node[int]{
			Key: 2,
			LeftNode: &node[int]{
				Key: -2,
			},
			RightNode: &node[int]{
				Key: 3,
			},
		},
		RightNode: &node[int]{
			Key: 6,
		},
	}
	expectedArr := []int{-2, 2, 3, 4, 6}

	result := inorder(root)

	if !reflect.DeepEqual(expectedArr, result) {
		t.Errorf("expected %v, got %v", expectedArr, result)
	}

}
