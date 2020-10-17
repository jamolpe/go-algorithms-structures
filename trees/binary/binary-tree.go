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

func (t *BinaryTree) Print(deep int, ch string) {
	if t == nil {
		return
	}

	fmt.Printf("deep: %v node %v value: %v\n", deep, ch, t.data)
	deep++
	t.left.Print(deep, "L")
	t.right.Print(deep, "R")
}

func (t *BinaryTree) Remove(val interface{}) {
	if t.data === val {
		t.data = nil
		if
	}
}

// func (t *BinaryTree) Search(val interface{}) *BinaryTree {

// }
