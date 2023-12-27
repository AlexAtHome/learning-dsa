package heap_test

import (
	"slices"
	"testing"

	"alexalex.dev/dsa/heap"
)

func TestBinHeapCreate(t *testing.T) {
	arr := []int8{3, 9, 2, 1, 4, 5}
	expected := []int8{9, 4, 5, 1, 3, 2}

	actual := heap.CreateBinHeap(slices.Clone(arr))

	if slices.Compare(actual.List, expected) != 0 {
		t.Fatalf("Expected %v to equal %v from %v!", actual, expected, arr)
	}
}

func TestBinHeapInsert(t *testing.T) {
	arr := []int8{9, 4, 5, 1, 3, 2}
	expected := []int8{9, 4, 7, 1, 3, 2, 5}

	h := heap.CreateBinHeap(slices.Clone(arr))
	h.Insert(7)

	if slices.Compare(h.List, expected) != 0 {
		t.Fatalf("Expected %v to equal %v from %v with 7 inserted!", h, expected, arr)
	}
}

func TestBinHeapDelete(t *testing.T) {
	arr := []int8{9, 3, 7, 1, 4, 2, 5}
	expected := []int8{9, 5, 7, 1, 4, 2}

	actual := heap.CreateBinHeap(slices.Clone(arr))
	actual.Delete(3)

	if slices.Compare(actual.List, expected) != 0 {
		t.Fatalf("%v hasn't been heapified properly from %v after removing 3! Expected %v", actual, arr, expected)
	}
}

func TestBinHeapPeek(t *testing.T) {
	arr := []int8{3, 9, 2, 1, 4, 5}
	expected := int8(9)

	h := heap.CreateBinHeap(slices.Clone(arr))

	if h.Peek() != expected {
		t.Fatalf("Expected the top item (%v) in %v to equal %v!", h.Peek(), arr, expected)
	}
}

func TestBinHeapExtract(t *testing.T) {
	arr := []int8{3, 9, 2, 1, 4, 5}
	expected := int8(9)

	h := heap.CreateBinHeap(slices.Clone(arr))
	extract := h.Extract()

	if extract != expected {
		t.Fatalf("Expected to extract %v from heaped %v!", expected, arr)
	}
	if slices.Contains(h.List, 9) {
		t.Fatalf("The top item (%v) did not actually extract from %v!", 9, h)
	}
}
