package deque

import (
	"fmt"
)

type Deque struct {
	list              []int
	size, front, rear int
}

func CreateDeque(size int) (dq *Deque) {
	dq = &Deque{make([]int, size, size), size, -1, 0}
	return dq
}

func (dq *Deque) GetFront() int {
	if dq.IsEmpty() {
		return -1
	}
	return dq.list[dq.front]
}

func (dq *Deque) GetRear() int {
	if dq.IsEmpty() {
		return -1
	}
	return dq.list[dq.rear]
}

func (dq *Deque) InsertFront(item int) error {
	if dq.IsFull() {
		return fmt.Errorf("Unable to insert into a full deque. %v", dq)
	}
	if dq.IsEmpty() {
		dq.front = 0
		dq.rear = 0
	} else if dq.front == 0 {
		dq.front = dq.size - 1
	} else {
		dq.front--
	}
	dq.list[dq.front] = item
	return nil
}

func (dq *Deque) InsertRear(item int) error {
	if dq.IsFull() {
		return fmt.Errorf("Unable to insert into a full deque. %v", dq)
	}
	if dq.IsEmpty() {
		dq.rear = 0
		dq.front = 0
	} else if dq.rear == dq.size-1 {
		dq.rear = 0
	} else {
		dq.rear++
	}
	dq.list[dq.rear] = item
	return nil
}

func (dq *Deque) DeleteFront() error {
	if dq.IsEmpty() {
		return fmt.Errorf("Unable to delete from an empty deque. %v", dq)
	}
	dq.list[dq.front] = 0
	if dq.front == dq.rear {
		dq.front = -1
		dq.rear = -1
	} else if dq.front == dq.size-1 {
		dq.front = 0
	} else {
		dq.front++
	}
	return nil
}

func (dq *Deque) DeleteRear() error {
	if dq.IsEmpty() {
		return fmt.Errorf("Unable to delete from an empty deque. %v", dq)
	}
	dq.list[dq.rear] = 0
	if dq.rear == dq.front {
		dq.front = -1
		dq.rear = -1
	} else if dq.rear == 0 {
		dq.rear = dq.size - 1
	} else {
		dq.rear--
	}
	return nil
}

func (dq *Deque) IsFull() bool {
	return (dq.front == 0 && dq.rear == dq.size-1) || dq.front == dq.rear+1
}

func (dq *Deque) IsEmpty() bool {
	return dq.front == -1
}

func (dq *Deque) String() string {
	return fmt.Sprintf("Deque: %v", dq.list)
}
