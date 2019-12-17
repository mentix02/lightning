// Commonly used data structures optimized for
// internal use cases concerning reading rows
// from MySQL databases. Also contains commonly
// used JSON response structures.
package main

// -------- JSON Responses --------- //

type DetailResponse struct {
	Detail string `json:"detail"`
}

// -------- Data Structures -------- //

// Singly Linked List //
type Node struct {
	next *Node
	value uint32
}

type List struct {
	len uint32
	head *Node
	tail *Node
}

// Adds an element to the list.
func (l *List) append(value uint32) {

	// Create node with empty value.
	node := Node{value:value}

	// If length of l is 0, then head
	// and tail are the same node.
	if l.len == 0 {
		l.head = &node
		l.tail = &node
	} else {
		// If there are some elements in list
		// then make the tail point to the new
		// node and change the tail to said node.
		tail := l.tail
		tail.next = &node
		l.tail = &node
	}

	// Increment length.
	l.len++
}

// Converts node to slice for easier usage
// and traversal.
func (l List) toSlice() []uint32 {
	s := make([]uint32, l.len)
	for e, i := l.head, 0; e != nil; e, i = e.next, i + 1 {
		s[i] = e.value
	}
	return s
}

// Singly Linked List //
