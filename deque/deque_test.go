package deque_test

import (
	"testing"

	"alexalex.dev/dsa/deque"
)

func TestDequeInsertRear(t *testing.T) {
	dq := deque.CreateDeque(5)
	dq.InsertRear(7)
	dq.InsertRear(3)
	dq.InsertRear(1)
	actual := dq.GetRear()
	if actual != 1 {
		t.Fatalf("Expected the rear value %v from %v to equal 1", actual, dq)
	}
}

func TestDequeInsertRearReverse(t *testing.T) {
	dq := deque.CreateDeque(5)
	dq.InsertFront(7)
	dq.InsertFront(3)
	dq.InsertFront(1)
	dq.DeleteRear()
	dq.DeleteRear()
	actual := dq.GetRear()
	if actual != 1 {
		t.Fatalf("Expected the rear value %v from %v to equal 1", actual, dq)
	}
}

func TestDequeInsertFront(t *testing.T) {
	dq := deque.CreateDeque(5)
	dq.InsertFront(7)
	dq.InsertFront(3)
	dq.InsertFront(1)
	actual := dq.GetFront()
	if actual != 1 {
		t.Fatalf("Expected the rear value %v from %v to equal 1", actual, dq)
	}
}

func TestDequeInsertFrontReverse(t *testing.T) {
	dq := deque.CreateDeque(5)
	dq.InsertRear(7)
	dq.InsertRear(3)
	dq.InsertRear(1)
	dq.InsertFront(5)
	actual := dq.GetFront()
	if actual != 5 {
		t.Fatalf("Expected the rear value %v from %v to equal 5", actual, dq)
	}
}

func TestDequeDeleteFront(t *testing.T) {
	dq := deque.CreateDeque(5)
	dq.InsertFront(7)
	dq.InsertFront(3)
	dq.InsertFront(1)
	dq.DeleteFront()
	actual := dq.GetFront()
	if actual != 3 {
		t.Fatalf("Expected the front value %v from %v to equal 3", actual, dq)
	}
}
