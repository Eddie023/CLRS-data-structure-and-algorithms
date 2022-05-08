package queue

type Queue struct {
	head   int
	tail   int
	values []int
}

func NewQueue() *Queue {
	return &Queue{0, 0, []int{}}
}

// Add element to a queue.
func (q *Queue) Enqueue(elm int) *Queue {
	q.values = append(q.values, elm)
	q.tail += 1

	return q
}

// Remove and return the first element of the queue.
func (q *Queue) Dequeue() any {
	elm := q.values[q.head]
	q.values = q.values[1:]
	q.head += 1

	return elm
}

func (q *Queue) ListItems() []int {
	return q.values
}
