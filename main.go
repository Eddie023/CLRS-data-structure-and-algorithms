package main

import (
	"fmt"

	"github.com/eddie023/data-structure/queue"
)

func main() {
	q := queue.NewQueue()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	q.Dequeue()

	fmt.Printf("the current value in queue %v is %d\n", q, q.ListItems())
}
