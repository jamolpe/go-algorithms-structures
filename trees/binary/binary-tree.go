package binaryTree

import (
	"fmt"
)

type BinaryTree struct {
	data          interface{}
	left          *BinaryTree
	right         *BinaryTree
	ComparatorMax func(interface{}, interface{}) bool
	Equal         func(interface{}, interface{}) bool
}

func New(comparable func(interface{}, interface{}) bool, equalable func(interface{}, interface{}) bool) *BinaryTree {
	return &BinaryTree{
		ComparatorMax: comparable,
		data:          nil,
		Equal:         equalable,
	}
}

func (t *BinaryTree) Insert(val interface{}) {
	if t.data == nil {
		t.data = val
		t.right = New(t.ComparatorMax, t.Equal)
		t.left = New(t.ComparatorMax, t.Equal)
	} else {
		if t.ComparatorMax(val, t.data) {
			t.right.Insert(val)
		} else {
			t.left.Insert(val)
		}
	}
}

func FindMinFromNode(pretree **BinaryTree) interface{} {
	tree := *pretree
	if tree.left.data == nil {
		value := tree.data
		if tree.right.data != nil {
			*pretree = tree.right
		} else {
			tree.data = nil
		}
		return value
	} else {
		return FindMinFromNode(&tree.left)
	}
}

func (t *BinaryTree) Search(val interface{}) interface{} {
	if t.data == nil {
		return nil
	}
	if t.Equal(t.data, val) {
		return t.data
	} else {
		if t.ComparatorMax(val, t.data) {
			return t.right.Search(val)
		} else {
			return t.left.Search(val)
		}
	}
}

func (t *BinaryTree) Remove(val interface{}) {
	if t.data != nil {
		if t.Equal(t.data, val) {
			if t.right.data == nil && t.left.data == nil {
				t.data = nil
			} else if t.left.data != nil && t.right.data == nil {
				t.data = t.left.data
			} else if t.right.data != nil && t.left.data == nil {
				t.data = t.right.data
			} else {
				t.data = FindMinFromNode(&t.right)
			}
		} else {
			if t.ComparatorMax(val, t.data) {
				t.right.Remove(val)
			} else {
				t.left.Remove(val)
			}
		}
	}
}

func (t *BinaryTree) Print(deep int, ch string) {
	if t.data == nil {
		return
	}

	fmt.Printf("deep: %v node %v value: %v\n", deep, ch, t.data)
	deep++
	t.left.Print(deep, "L")
	t.right.Print(deep, "R")
}
