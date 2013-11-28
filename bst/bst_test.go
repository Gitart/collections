package bst

import "testing"

func TestInsert(t *testing.T) {
	expected := []int{5, 3, 7, 4, 6}
	bst := new(BST)

	for _, i := range expected {
		if !bst.Insert(i, i) {
			t.Errorf("Element %v should have been added to the tree", i)
		}
	}

	for _, i := range expected {
		if bst.Find(i) == nil {
			t.Errorf("Element %v expected to be in the tree, but was not", i)
		}
	}

	if bst.Insert(4, 44) {
		t.Error("Duplicate elements should not be added")
	}

	if bst.Find(4) == 44 {
		t.Error("Previously inserted elements should not be updated")
	}

	if c := bst.c; c != len(expected) {
		t.Errorf("Tree expected to have %v elements, but has %v instead", len(expected), c)
	}
}

func TestRemove(t *testing.T) {
	expected := []int{5, 3, 7, 4, 6}
	bst := new(BST)

	for _, i := range expected {
		bst.Insert(i, i)
	}

	if !bst.Delete(6) {
		t.Errorf("Element %v should have been removed from the tree", 6)
	}

	for _, i := range expected[0:3] {
		if bst.Find(i) == nil {
			t.Errorf("Element %v expected to be in the tree, but was not", i)
		}
	}

	if d := expected[len(expected)-1]; bst.Find(d) != nil {
		t.Errorf("Element %v should have been removed from the tree", d)
	}

	if bst.Delete(6) {
		t.Error("Duplicate elements should not be delete")
	}

	if c := bst.c; c != len(expected)-1 {
		t.Errorf("Tree expected to have %v elements, but has %v instead", len(expected)-1, c)
	}
}

func TestTraverse_InOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{3, 4, 5, 6, 7}

	bst := new(BST)

	for _, i := range elements {
		bst.Insert(i, i)
	}

	i := 0
	for e := range bst.Traverse(InOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestTraverse_PreOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{5, 3, 4, 7, 6}

	bst := new(BST)

	for _, i := range elements {
		bst.Insert(i, i)
	}

	i := 0
	for e := range bst.Traverse(PreOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestTraverse_PostOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{4, 3, 6, 7, 5}

	bst := new(BST)

	for _, i := range elements {
		bst.Insert(i, i)
	}

	i := 0
	for e := range bst.Traverse(PostOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}
