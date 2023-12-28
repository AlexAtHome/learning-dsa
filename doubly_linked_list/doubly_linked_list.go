package doublylinkedlist

import "fmt"

type Node struct {
	Value      string
	Next, Prev *Node
}

func (node Node) String() string {
	if node.Next != nil {
		return fmt.Sprintf("[Node: %v. Next: %v]", node.Value, node.Next.Value)
	} else {
		return fmt.Sprintf("[Node: %v]", node.Value)
	}
}

type DoublyLinkedList struct {
	head *Node
}

func (list *DoublyLinkedList) GetHead() *Node {
	return list.head
}

func (list *DoublyLinkedList) InsertHead(value string) (node *Node) {
	node = &Node{value, nil, nil}
	if list.head == nil {
		list.head = node
		list.head.Prev = nil
		list.head.Next = nil
	} else {
		node.Next = list.head
		node.Prev = nil
		list.head = node
	}
	return node
}

func (list *DoublyLinkedList) Insert(value string) *Node {
	if list.head == nil {
		return list.InsertHead(value)
	}
	last := list.head
	for last.Next != nil {
		last = last.Next
	}
	node := &Node{value, nil, last}
	last.Next = node

	return node
}

func (list *DoublyLinkedList) InsertAfter(prevNode *Node, value string) (newItem *Node) {
	if prevNode == nil {
		return nil
	}
	newItem = &Node{value, prevNode.Next, prevNode}
	prevNode.Next = newItem
	return newItem
}

func (list *DoublyLinkedList) InsertAt(position uint, value string) (node *Node) {
	prevItem := list.FindByIndex(int(position) - 1)
	if prevItem == nil {
		return nil
	}
	node = &Node{value, prevItem.Next, prevItem}
	prevItem.Next = node

	return node
}

func (list *DoublyLinkedList) Find(value string) *Node {
	current := list.head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (list *DoublyLinkedList) FindByIndex(index int) (item *Node) {
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

func (list *DoublyLinkedList) DeleteHead() {
	if list.head == nil {
		return
	}
	newHead := list.head.Next
	list.head.Next = nil
	list.head = newHead
}

func (list *DoublyLinkedList) DeleteTail() {
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

func (list *DoublyLinkedList) DeleteAt(index uint) {
	if index == 0 {
		list.DeleteHead()
	}
	prevItem := list.FindByIndex(int(index) - 1)
	if prevItem == nil || prevItem.Next == nil {
		return
	}
	prevItem.Next = prevItem.Next.Next
}

func (list *DoublyLinkedList) BubbleSort() {
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

func (list DoublyLinkedList) String() string {
	result := fmt.Sprintf("[%v", list.head.Value)
	temp := list.head.Next
	for temp != nil {
		result += fmt.Sprintf(" --> ")
		result += fmt.Sprintf("%v", temp.Value)
		temp = temp.Next
	}
	result += fmt.Sprintf("]")
	return result
}
