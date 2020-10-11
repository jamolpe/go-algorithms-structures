package binaryTree

type BinaryTree struct {
	data       interface{}
	left       *BinaryTree
	right      *BinaryTree
	Comparable func(interface{}, interface{}) bool
}

func New(comparable func(interface{}, interface{}) bool) *BinaryTree {
	return &BinaryTree{
		Comparable: comparable,
		data:       nil,
	}
}

func (t *BinaryTree) Insert(interface{}) {

}

func (t *BinaryTree) Remove(interface{}) {

}

func (t *BinaryTree) Search(interface{}) {

}
