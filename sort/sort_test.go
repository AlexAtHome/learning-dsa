package sort_test

import (
	"testing"

	"alexalex.dev/dsa/sort"
)

func isEqual(a []int, b []int) bool {
	if len(a) < len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestShouldSortAnArray(t *testing.T) {
	expected := []int{0, 2, 4, 5, 6, 7, 8, 9}
	foo := []int{8, 4, 5, 2, 9, 6, 0, 7}
	sort.Bubble(foo)
	if isEqual(foo, expected) == false {
		t.Fatalf("Failed to bubble sort the array %v", foo)
	}
}

func TestShouldReturnAnArgumentArrayWithOneElement(t *testing.T) {
	foo := []int{111}
	sort.Bubble(foo)
	if isEqual([]int{111}, foo) == false {
		t.Fatalf("The array %v was not supposed to change", foo)
	}
}
