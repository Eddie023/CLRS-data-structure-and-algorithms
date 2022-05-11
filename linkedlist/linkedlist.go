package linkedlist

import "fmt"

type LinkedList struct {
	head *Node
	len  int
}

type Node struct {
	value any
	next  *Node
}

// Create a new linkedlist with head pointing to empty nil.
func New() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

func (l *LinkedList) Size() int {
	return l.len
}

func (l *LinkedList) isEmpty() bool {
	return (l.len == 0)
}

// Insert at the start of the linkedlist.
func (l *LinkedList) Insert(value any) {

	// create new node with the value.
	newNode := Node{
		value: value,
	}

	// if empty, head will now point to the new node.
	if l.isEmpty() {
		l.head = &newNode
		l.len++

		return
	}

	// else, new node will now point to the node previously pointed by head.
	newNode.next = l.head

	// head will now point to new node.
	l.head = &newNode
	l.len++
}

func (l *LinkedList) ListItems() {
	tmp := l.head
	for tmp != nil {
		fmt.Println("value is", tmp.value)
		tmp = tmp.next
	}
}
