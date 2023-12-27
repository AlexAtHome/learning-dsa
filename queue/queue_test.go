package queue_test

import (
	"testing"

	"alexalex.dev/dsa/queue"
)

func TestCreateQueueEmpty(t *testing.T) {
	q := queue.CreateQueue(5)
	actual := q.IsEmpty()
	if actual != true {
		t.Fatal("Created queue is not empty!")
	}
}

func TestEnqueue(t *testing.T) {
	q := queue.CreateQueue(3)
	q.Enqueue("foo")
	isEmpty := q.IsEmpty()
	isFull := q.IsFull()
	if isEmpty {
		t.Fatalf("The queue %v doesn't count as non-empty", q)
	}
	if isFull {
		t.Fatalf("The queue %v doesn't count as non-full", q)
	}
}

func TestDequeue(t *testing.T) {
	q := queue.CreateQueue(3)
	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("baz")
	actual := q.Dequeue()
	if actual != "foo" {
		t.Fatalf("Expected 'foo' to get dequeue'd. Got %v instead!", actual)
	}
}

func TestQueueIsFull(t *testing.T) {
	q := queue.CreateQueue(3)
	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("baz")
	actual := q.IsFull()
	if actual != true {
		t.Fatalf("The queue %v doesn't count as full!", q)
	}
}

func TestQueueSize(t *testing.T) {
	q := queue.CreateQueue(3)
	q.Enqueue("foo")
	q.Enqueue("bar")
	actual := q.GetSize()
	if actual != 2 {
		t.Fatalf("Expected the queue %v to have size of 2, not %v!", q, actual)
	}
}
