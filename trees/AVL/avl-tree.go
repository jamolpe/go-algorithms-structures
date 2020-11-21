package AVL

import "fmt"

type AVLTree struct {
	data          interface{}
	left          *AVLTree
	right         *AVLTree
	ComparatorMax func(interface{}, interface{}) bool
	Equal         func(interface{}, interface{}) bool
	height        int
}

func New(comparable func(interface{}, interface{}) bool, equalable func(interface{}, interface{}) bool) *AVLTree {
	return &AVLTree{
		ComparatorMax: comparable,
		data:          nil,
		Equal:         equalable,
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *AVLTree) calculateHeight() {
	t.height = 1 + max(t.right.height, t.left.height)
}

func (t *AVLTree) rotateLeft() {

}

func (t *AVLTree) rotateRigth() {

}

func (t *AVLTree) balanceTree() *AVLTree {
	balance := calcleftH(t.left) - calcrightH(t.right)
	if balance == -2 {
		// rotate left

	} else if balance == 2 {
		// rotate rigth
	}

}

func (t *AVLTree) IsBalanced() bool {
	rightBal, leftBal := calcrightH(t.right), calcleftH(t.left)
	if (leftBal - rightBal) > 1 {
		return false
	} else {
		return true
	}
}

func (t *AVLTree) Insert(val interface{}) {
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

func (t *AVLTree) Print(deep int, ch string) {
	if t.data == nil {
		return
	}

	fmt.Printf("deep: %v node %v value: %v\n", deep, ch, t.data)
	deep++
	t.left.Print(deep, "L")
	t.right.Print(deep, "R")
}
