package circularqueue

import "fmt"

type CircularQueue struct {
	list              []string
	size, front, rear int8
}

func CreateCircularQueue(size int8) *CircularQueue {
	list := make([]string, size, size)
	return &CircularQueue{list, size, -1, -1}
}

func (cq *CircularQueue) Enqueue(item string) error {
	if cq.IsFull() {
		return fmt.Errorf("Cant enqueue any more items in the full circular queue!")
	}
	if cq.IsEmpty() {
		cq.front = 0
	}
	cq.rear = (cq.rear + 1) % cq.size
	cq.list[cq.rear] = item
	return nil
}

func (cq *CircularQueue) Dequeue() (item string) {
	if cq.IsEmpty() {
		return ""
	}
	item = cq.list[cq.front]
	cq.list[cq.front] = ""
	if cq.front == cq.rear {
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % cq.size
	}
	return item
}

func (cq *CircularQueue) IsFull() bool {
	return (cq.front == 0 && cq.rear == cq.size-1) || (cq.front == cq.rear+1)
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.front == -1
}

func (cq *CircularQueue) String() string {
	if cq.IsEmpty() {
		return fmt.Sprintf("Empty circular queue")
	}
	return fmt.Sprintf("Circular Queue: %v", cq.list)
}
