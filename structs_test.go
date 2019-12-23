package main

import "testing"

var (
	elements = []uint32{3, 2, 6}
)

func equal(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func generateList() List {
	l := List{}
	for j := range elements {
		l.append(elements[j])
	}
	return l
}

func TestListTraversal(t *testing.T) {
	l := generateList()
	for e, i := l.head, 0; e != nil; e, i = e.next, i+1 {
		if e.value == elements[i] {
			continue
		} else {
			t.Errorf("e.value != %d", elements[i])
		}
	}
}

func TestListConversionToSlice(t *testing.T) {
	l := generateList()
	s := l.toSlice()
	if !equal(s, elements) {
		t.Errorf("%v != %v", s, elements)
	}
}
