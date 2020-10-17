package main

import "fmt"

func comparable(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

func main() {
	createExample()
	removeExample()
}

func removeExample() {
	fmt.Println("remove example")
	tree := New(comparable)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(11)
	tree.Insert(15)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Remove(6)
	tree.Print(0, "M")
}
func createExample() {
	fmt.Println("create example")
	tree := New(comparable)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(2)
	tree.Print(0, "M")
}
