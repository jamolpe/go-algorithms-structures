package main

import (
	"fmt"
)

type BinaryTree struct {
	data          interface{}
	left          *BinaryTree
	right         *BinaryTree
	ComparatorMax func(interface{}, interface{}) bool
}

func New(comparable func(interface{}, interface{}) bool) *BinaryTree {
	return &BinaryTree{
		ComparatorMax: comparable,
		data:          nil,
	}
}

func (t *BinaryTree) Insert(val interface{}) {
	if t.data == nil {
		t.data = val
		t.right = New(t.ComparatorMax)
		t.left = New(t.ComparatorMax)
	} else {
		if t.ComparatorMax(val, t.data) {
			t.right.Insert(val)
		} else {
			t.left.Insert(val)
		}
	}
}

func findMinFromNode(pretree **BinaryTree) interface{} {
	tree := *pretree
	if tree.left.data == nil {
		value := tree.data
		fmt.Printf("found value %v \n", value)
		if tree.right.data != nil {
			*pretree = tree.right
		} else {
			tree.data = nil
		}
		return value
	} else {
		fmt.Printf("node value: %v, leftValue: %v not found next one\n", tree.data, tree.left.data)
		return findMinFromNode(&tree.left)
	}
}

func (t *BinaryTree) Remove(val interface{}) {
	if val == t.data {
		if t.right == nil && t.left == nil {
			t.data = nil
		} else if t.left != nil && t.right == nil {
			t.data = t.left.data
		} else if t.right != nil && t.left == nil {
			t.data = t.right.data
		} else {
			t.data = findMinFromNode(&t.right)
		}
	}
}

func (t *BinaryTree) Print(deep int, ch string) {
	if t == nil {
		return
	}

	fmt.Printf("deep: %v node %v value: %v\n", deep, ch, t.data)
	deep++
	t.left.Print(deep, "L")
	t.right.Print(deep, "R")
}
