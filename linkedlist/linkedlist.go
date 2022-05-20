package linkedlist

type LinkedList struct {
	head *Node
	len  int
}

type Node struct {
	value int
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

func createNewNode(v int) Node {
	return Node{
		value: v,
	}
}

// Insert at the start of the linkedlist.
func (l *LinkedList) Insert(value int) {

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

func (l *LinkedList) InsertAtLast(value int) {

	newNode := createNewNode(value)

	tmp := l.head
	// loop until the end node
	for tmp.next != nil {
		// fmt.Println("The value is", tmp.value)
		tmp = tmp.next
	}

	tmp.next = &newNode
	newNode.next = nil
}

func (l *LinkedList) GetLastItem() int {

	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}

	return tmp.value
}

func (l *LinkedList) ListItems() []int {
	tmp := l.head

	items := []int{}
	for tmp != nil {
		// fmt.Println("value is", tmp.value)
		items = append(items, tmp.value)
		tmp = tmp.next
	}

	return items
}
