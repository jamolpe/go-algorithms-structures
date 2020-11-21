package binaryTree

import (
	"testing"
)

func comparable(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

func equalable(a interface{}, b interface{}) bool {
	return a.(int) == b.(int)
}

func Test_removeExample(t *testing.T) {
	tree := New(comparable, equalable)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(11)
	tree.Insert(15)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Remove(3)
	tree.Print(0, "M")
	min := FindMinFromNode(&tree)
	if min != 4 {
		t.Errorf("removed wrong values")
	}
}
func Test_removeNotExistingExample(t *testing.T) {
	tree := New(comparable, equalable)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(11)
	tree.Insert(15)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Remove(1)
	tree.Print(0, "M")
	min := FindMinFromNode(&tree)
	if min != 3 {
		t.Errorf("removed wrong values")
	}
}

func Test_createExample(t *testing.T) {
	tree := New(comparable, equalable)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(2)
	min := FindMinFromNode(&tree)
	if min != 2 {
		t.Errorf("insert is not working fine")
	}
}

type Searchable struct {
	Key         int
	StoredValue string
}

func comparableForSearch(a interface{}, b interface{}) bool {
	return a.(Searchable).Key > b.(Searchable).Key
}
func equalableForSearch(a interface{}, b interface{}) bool {
	return a.(Searchable).Key == b.(Searchable).Key
}

func Test_searchExample(t *testing.T) {
	tree := New(comparableForSearch, equalableForSearch)
	tree.Insert(Searchable{Key: 6, StoredValue: "pep 6"})
	tree.Insert(Searchable{Key: 2, StoredValue: "prp 3"})
	tree.Insert(Searchable{Key: 1, StoredValue: "prp 1"})
	tree.Insert(Searchable{Key: 11, StoredValue: "prp 11"})
	tree.Insert(Searchable{Key: 15, StoredValue: "prp 15"})
	tree.Insert(Searchable{Key: 8, StoredValue: "prp 8"})
	tree.Insert(Searchable{Key: 9, StoredValue: "prp 9"})
	tree.Insert(Searchable{Key: 10, StoredValue: "prp 10"})
	result := tree.Search(Searchable{Key: 6})
	if result.(Searchable).StoredValue != "pep 6" {
		t.Errorf("find value not correct")
	}
}

func Test_searchExampleNotExistent(t *testing.T) {
	tree := New(comparableForSearch, equalableForSearch)
	tree.Insert(Searchable{Key: 6, StoredValue: "pep 6"})
	tree.Insert(Searchable{Key: 2, StoredValue: "prp 3"})
	tree.Insert(Searchable{Key: 1, StoredValue: "prp 1"})
	tree.Insert(Searchable{Key: 11, StoredValue: "prp 11"})
	tree.Insert(Searchable{Key: 15, StoredValue: "prp 15"})
	tree.Insert(Searchable{Key: 8, StoredValue: "prp 8"})
	tree.Insert(Searchable{Key: 9, StoredValue: "prp 9"})
	tree.Insert(Searchable{Key: 10, StoredValue: "prp 10"})
	result := tree.Search(Searchable{Key: 99})
	if result != nil {
		t.Errorf("error is not null")
	}
}
