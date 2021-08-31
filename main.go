package main

import "fmt"

func main() {
	q := NewQueue()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	q.Dequeue()

	fmt.Println(q)
}
