package circularqueue_test

import (
	"testing"

	circularqueue "alexalex.dev/dsa/circular_queue"
)

func TestCircularQueueCreate(t *testing.T) {
	cq := circularqueue.CreateCircularQueue(8)
	actual := cq.IsEmpty()
	if actual != true {
		t.Fatalf("For some reason, a new circular queue %v doesn't count as empty!", cq)
	}
}

func TestCircularQueueEnqueue(t *testing.T) {
	cq := circularqueue.CreateCircularQueue(3)
	cq.Enqueue("foo")
	if cq.IsEmpty() {
		t.Fatalf("The queue %v isn't supposed to be empty!", cq)
	}
}

func TestCircularQueueDequeue(t *testing.T) {
	cq := circularqueue.CreateCircularQueue(3)
	cq.Enqueue("foo")
	cq.Enqueue("bar")
	actual := cq.Dequeue()
	if actual != "foo" {
		t.Fatalf("Expected to dequeue 'foo' instead of %v", actual)
	}
}

func TestCircularQueueReEnqueue(t *testing.T) {
	cq := circularqueue.CreateCircularQueue(3)
	cq.Enqueue("foo")
	cq.Enqueue("bar")
	cq.Enqueue("baz")
	cq.Dequeue()
	cq.Dequeue()
	cq.Enqueue("boo")
	cq.Enqueue("far")
	actual := cq.Dequeue()
	if actual != "baz" {
		t.Fatalf("Expected to dequeue 'baz' instead of %v", actual)
	}
}
