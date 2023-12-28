package linkedlist_test

import (
	"testing"

	linkedlist "alexalex.dev/dsa/linked_list"
)

func TestNodeNext(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	bar := list.Insert("bar")
	list.Insert("baz")
	actual := list.GetHead().Next
	if actual.Value != "bar" {
		t.Fatalf("Expected the current node to be %v, but got %v", bar, actual)
	}
}

func TestNodeNextNull(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	actual := list.GetHead().Next
	if actual != nil {
		t.Fatalf("Expected the current node in %v to be nil, but got %v", list, actual)
	}
}

func TestLinkedListInsertHead(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.InsertHead("foo")
	actual := list.GetHead()
	if actual == nil {
		t.Fatalf("Failed to insert the head into %v", list)
	}
}

func TestLinkedListInsertTail(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	bar := list.Insert("bar")
	list.Insert("baz")

	actual := list.GetHead().Next
	if actual.Value != "bar" {
		t.Fatalf("Expected the current node to be %v, but got %v", bar, actual)
	}
}

func TestLinkedListInsertAfter(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	bar := list.Insert("bar")
	list.Insert("baz")
	list.Insert("boo")
	far := list.InsertAfter(bar, "far")

	actual := bar.Next
	if actual.Value != far.Value {
		t.Fatalf("Expected %v to be inserted after %v in %v", far, bar, list)
	}
}

func TestLinkedListInsertAt(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	bar := list.Insert("bar")
	list.Insert("baz")
	list.Insert("boo")
	far := list.InsertAt(2, "far")

	actual := bar.Next
	if actual.Value != "far" {
		t.Fatalf("Expected %v next node to be %v, but got %v", bar, far, actual)
	}
}

func TestLinkedListDeleteHead(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	bar := list.Insert("bar")
	list.Insert("baz")
	list.Insert("boo")
	list.DeleteHead()

	head := list.GetHead()
	if head.Value != bar.Value {
		t.Fatalf("Expected %v to be %v as the head of %v", head, bar, list)
	}
}

func TestLinkedListDeleteTail(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	list.Insert("bar")
	baz := list.Insert("baz")
	list.Insert("boo")
	list.DeleteTail()

	current := list.GetHead().Next.Next
	if current.Value != baz.Value {
		t.Fatalf("Expected %v to be %v as the tail of %v", current, baz, list)
	}
}

func TestLinkedListDeleteAt(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	list.Insert("bar")
	baz := list.Insert("baz")
	list.DeleteAt(1)

	next := list.GetHead().Next
	if next.Value != baz.Value {
		t.Fatalf("Expected %v to be %v as the next value of %v", next, baz, list)
	}
}

func TestLinkedListDeleteAtHead(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	bar := list.Insert("bar")
	list.Insert("baz")
	list.DeleteAt(0)

	head := list.GetHead()
	if head.Value != bar.Value {
		t.Fatalf("Expected %v to be %v as the next value of %v", head, bar, list)
	}
}

func TestLinkedListFind(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	list.Insert("bar")
	baz := list.Insert("baz")
	found := list.Find("baz")

	if found.Value != baz.Value {
		t.Fatalf("Expected to find %v in %v instead of %v", baz, list, found)
	}
}

func TestLinkedListBubbleSort(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Insert("foo")
	list.Insert("bar")
	list.Insert("baz")
	list.Insert("boo")
	list.BubbleSort()

	head := list.GetHead()
	if head.Next.Value != "baz" {
		t.Fatalf("Expected 'baz' to be next to %v instead of %v in %v", head, head.Next, list)
	}
}
