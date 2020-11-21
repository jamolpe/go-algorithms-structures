package AVL

import "testing"

func comparable(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

func equalable(a interface{}, b interface{}) bool {
	return a.(int) == b.(int)
}

func Test_example(t *testing.T) {
	tree := New(comparable, equalable)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(11)
	tree.Insert(15)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Print(0, "M")
	isBalanced := tree.IsBalanced()
	println(isBalanced)
}
