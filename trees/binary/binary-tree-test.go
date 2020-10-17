package main

func comparable(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

func main() {
	tree := New(comparable)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(2)
	tree.Print(0, "M")
}
