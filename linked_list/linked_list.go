package linkedlist

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

func (node Node) String() string {
	if node.Next != nil {
		return fmt.Sprintf("[Node: %v. Next: %v]", node.Value, node.Next.Value)
	} else {
		return fmt.Sprintf("[Node: %v]", node.Value)
	}
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) GetHead() *Node {
	return list.head
}

func (list *LinkedList) InsertHead(value string) (node *Node) {
	node = &Node{value, nil}
	if list.head == nil {
		list.head = node
	} else {
		temp := list.head
		list.head = node
		list.head.Next = temp
	}
	return node
}

func (list *LinkedList) Insert(value string) (node *Node) {
	node = &Node{value, nil}
	if list.head == nil {
		list.head = node
	} else {
		last := list.head
		for last.Next != nil {
			last = last.Next
		}
		last.Next = node
	}
	return node
}

func (list *LinkedList) InsertAfter(prevNode *Node, value string) (newItem *Node) {
	if prevNode == nil {
		return nil
	}
	newItem = &Node{value, prevNode.Next}
	prevNode.Next = newItem
	return newItem
}

func (list *LinkedList) InsertAt(position uint, value string) (node *Node) {
	prevItem := list.FindByIndex(int(position) - 1)
	if prevItem == nil {
		return nil
	}
	node = &Node{value, prevItem.Next}
	prevItem.Next = node

	return node
}

func (list *LinkedList) Find(value string) *Node {
	current := list.head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (list *LinkedList) FindByIndex(index int) (item *Node) {
	if index < 0 {
		return nil
	}

	i := 0
	temp := list.head
	for i <= index || item == nil || temp != nil {
		if i == index {
			item = temp
		}
		temp = temp.Next
		i++
	}
	return item
}

func (list *LinkedList) DeleteHead() {
	if list.head == nil {
		return
	}
	newHead := list.head.Next
	list.head.Next = nil
	list.head = newHead
}

func (list *LinkedList) DeleteTail() {
	if list.head == nil {
		return
	}
	if list.head.Next == nil {
		list.head = nil
	}
	secondToLast := list.head
	for secondToLast.Next.Next != nil {
		secondToLast = secondToLast.Next
	}
	secondToLast.Next = nil
}

func (list *LinkedList) DeleteAt(index uint) {
	if index == 0 {
		list.DeleteHead()
	}
	prevItem := list.FindByIndex(int(index) - 1)
	if prevItem == nil || prevItem.Next == nil {
		return
	}
	prevItem.Next = prevItem.Next.Next
}

func (list *LinkedList) BubbleSort() {
	current := list.head
	for current != nil {
		index := current.Next
		for index != nil {
			if current.Value > index.Value {
				temp := current.Value
				current.Value = index.Value
				index.Value = temp
			}
			index = index.Next
		}
		current = current.Next
	}
}

func (list LinkedList) String() string {
	temp := list.head
	i := 0
	result := "["
	for temp != nil {
		if i > 0 {
			result += fmt.Sprintf(" --> ")
		}
		result += fmt.Sprintf("%v", temp.Value)
		temp = temp.Next
		i++
	}
	result += fmt.Sprintf("]")
	return result
}
