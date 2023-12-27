package stack_test

import (
	"testing"

	"alexalex.dev/dsa/stack"
)

func TestStackCreateEmpty(t *testing.T) {
	stack := stack.CreateStack(3)
	if !stack.IsEmpty() {
		t.Fatal("Stack is non empty!")
	}
}

func TestStackPush(t *testing.T) {
	stack := stack.CreateStack(2)
	stack.Push("foo")
	if stack.GetSize() != 1 {
		t.Fatal("An item 'foo' wasn't pushed correctly into the stack")
	}
}

func TestStackPushMaxSize(t *testing.T) {
	stack := stack.CreateStack(2)
	stack.Push("foo")
	stack.Push("bar")
	stack.Push("baz")
	if stack.GetSize() > 2 {
		t.Fatal("Stack maximum size exceeded!")
	}
}

func TestStackPushMultipleMaxSize(t *testing.T) {
	stack := stack.CreateStack(2)
	stack.Push("foo", "bar", "baz", "fab")
	if stack.GetSize() > 2 {
		t.Fatal("Stack maximum size exceeded!")
	}
}

func TestStackPop(t *testing.T) {
	stack := stack.CreateStack(2)
	stack.Push("foo", "bar")
	actual := stack.Pop()
	if actual != "bar" {
		t.Fatalf("The last item in stack is not popped properly! Expected bar instead of %v", actual)
	}
}

func TestStackPeek(t *testing.T) {
	stack := stack.CreateStack(2)
	stack.Push("foo")
	stack.Push("bar")
	actual := stack.Peek()
	if actual != "bar" {
		t.Fatalf("It peeked at the wrong item! Expected bar instead of %v", actual)
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := stack.CreateStack(2)
	stack.Push("foo")
	stack.Pop()
	actual := stack.IsEmpty()
	if actual != true {
		t.Fatal("Expected stack to be empty!")
	}
}

func TestStackIsFull(t *testing.T) {
	stack := stack.CreateStack(2)
	stack.Push("foo")
	stack.Push("bar")
	actual := stack.IsFull()
	if actual != true {
		t.Fatal("Expected stack to be full!")
	}
}

func TestStackGetSize(t *testing.T) {
	stack := stack.CreateStack(2)
	stack.Push("foo")
	actual := stack.GetSize()
	if actual != 1 {
		t.Fatalf("Expected stack size to be 1, not %v!", actual)
	}
}
