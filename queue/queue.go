package queue

import (
	"fmt"

	"golang.org/x/exp/maps"
)

// A simple queue that follows the FIFO rule
type Queue struct {
	MaxSize int
	front   int
	rear    int
	list    map[int]string
}

func CreateQueue(size int) *Queue {
	return &Queue{size, -1, -1, map[int]string{}}
}

func (q *Queue) Enqueue(item string) error {
	if q.IsFull() {
		return fmt.Errorf("Cant enqueue any more items in the full queue!")
	}
	if q.front < 0 {
		q.front = 0
	}
	q.rear++
	q.list[q.rear] = item
	return nil
}

func (q *Queue) Dequeue() (item string) {
	if q.IsEmpty() {
		return ""
	}
	item = q.list[q.front]
	delete(q.list, q.front)
	if q.GetSize() == 0 {
		q.front = -1
		q.rear = -1
	} else {
		q.front++
	}
	return item
}

func (q *Queue) IsEmpty() bool {
	return q.front == -1
}

func (q *Queue) IsFull() bool {
	return q.front > -1 && q.rear == q.MaxSize-1
}

func (q *Queue) GetSize() int {
	if q.front == -1 {
		return 0
	}
	if q.front == q.rear {
		return 1
	}
	return q.rear - q.front + 1
}

func (q *Queue) String() string {
	return fmt.Sprintf("Queue: %v", maps.Values(q.list))
}
