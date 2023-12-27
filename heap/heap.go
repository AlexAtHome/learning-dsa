package heap

import (
	"fmt"
	"slices"
)

type Heap struct {
	List []int8
}

func CreateBinHeap(arr []int8) (heapTree *Heap) {
	heapTree = &Heap{arr}
	heapTree.Heapify()
	return heapTree
}

func (heap *Heap) Heapify() {
	size := len(heap.List)
	for i := size/2 - 1; i >= 0; i-- {
		heap.heapifyNode(i)
	}
}

func (heap *Heap) heapifyNode(i int) {
	largest := i
	size := len(heap.List)
	leftChild := 2*i + 1
	rightChild := 2*i + 2

	if leftChild < size && heap.List[leftChild] > heap.List[largest] {
		largest = leftChild
	} else if rightChild < size && heap.List[rightChild] > heap.List[largest] {
		largest = rightChild
	}

	if largest != i {
		Swap(heap.List, largest, i)
		heap.heapifyNode(largest)
	}
}

func (heap *Heap) Insert(item int8) {
	heap.List = append(heap.List, item)
	if len(heap.List) > 1 {
		heap.Heapify()
	}
}

func (heap *Heap) Delete(item int8) {
	index := slices.Index(heap.List, item)
	if index == -1 {
		return
	}
	size := len(heap.List)
	Swap(heap.List, index, size-1)
	heap.List = heap.List[:size-1]
	heap.Heapify()
}

func (heap *Heap) Peek() int8 {
	return heap.List[0]
}

func (heap *Heap) Extract() (item int8) {
	defer func() {
		heap.List = heap.List[1:]
		heap.Heapify()
	}()
	return heap.List[0]
}

func (heap *Heap) String() string {
	if len(heap.List) == 0 {
		return "[An empty heap]"
	}
	return fmt.Sprintf("Heap: %v", heap.List)
}

// A misc function for swapping the elements by their indexes.
func Swap[K any](arr []K, left int, right int) {
	temp := arr[left]
	arr[left] = arr[right]
	arr[right] = temp
}
