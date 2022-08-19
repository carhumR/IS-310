package main

import (
	"errors"
	"fmt"
)

type PriorityQueue struct {
	High []int
	Low  []int
}

func (q *PriorityQueue) Enqueue(elem int, highPriority bool) {
	if highPriority {
		q.High = append(q.High, elem)
	} else {
		q.Low = append(q.Low, elem)
	}
}

func (q *PriorityQueue) Dequeue() (int, error) {

	if len(q.High) != 0 {
		var firstElement int
		firstElement, q.High = q.High[0], q.High[1:]
		return firstElement, nil
	}

	if len(q.Low) != 0 {
		var firstElement int
		firstElement, q.Low = q.Low[0], q.Low[1:]
		return firstElement, nil
	}

	return 0, errors.New("empty queue")
}

func (q *PriorityQueue) Length() int {
	return len(q.High) + len(q.Low)
}

func (q *PriorityQueue) Peek() (int, error) {
	if len(q.High) != 0 {
		return q.High[0], nil
	}
	if len(q.Low) != 0 {
		return q.Low[0], nil
	}
	return 0, errors.New("empty queue")
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.Length() == 0
}

func main() {
	fmt.Println("Queues Section")

	queue := PriorityQueue{}

	fmt.Println(queue)
	queue.Enqueue(1, true)
	fmt.Println(queue)
	queue.Enqueue(10, false)
	fmt.Println(queue)

	elem, _ := queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)

	queue.Enqueue(2, true)
	fmt.Println(queue)

	elem, _ = queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)

	elem, _ = queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)

}