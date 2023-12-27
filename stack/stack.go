package stack

import (
	"fmt"
)

type Stack struct {
	list    []string
	top     int
	maxSize int
}

func CreateStack(maxSize int) *Stack {
	return &Stack{[]string{}, -1, maxSize}
}

func (s *Stack) Push(items ...string) error {
	if s.IsFull() {
		return fmt.Errorf("Cant push any more items in the full stack!")
	}
	length := len(items)
	if (s.top + 1 + length) > s.maxSize {
		return fmt.Errorf("Too many items for a stack!")
	}
	s.top = s.top + length
	s.list = append(s.list, items...)
	return nil
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	item := s.list[s.top]
	if s.GetSize() <= 1 {
		s.list = []string{}
		s.top = -1
	} else {
		s.list = append(s.list[:s.top], s.list[s.top+1:]...)
		s.top = s.top - 1
	}
	return item
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return s.list[s.top]
}

func (s *Stack) IsFull() bool {
	return s.GetSize() == s.maxSize
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) GetSize() int {
	return s.top + 1
}

func StackTest() {
	stack := CreateStack(3)
	stack.Push("foo")
	stack.Push("bar")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Push("foo")
	stack.Push("bar")
	stack.Push("baz")
	err := stack.Push("biz")
	if err != nil {
		fmt.Println(err)
	}
}
