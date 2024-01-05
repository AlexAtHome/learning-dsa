package hashtable_test

import (
	"testing"

	ht "alexalex.dev/dsa/hash_table"
)

func TestHashTableInsert(t *testing.T) {
	table := ht.CreateHashTable(10)
	table.Insert(111, "foo")
	table.Insert(113, "bar")
	table.Insert(123, "baz")
	size := table.Size
	if size != 3 {
		t.Fatalf("Expected the size of hash table of 3, got %v", size)
	}
}

func TestHashTableHashCollision(t *testing.T) {
	table := ht.CreateHashTable(1000)
	table.Insert(10, "foo")
	table.Insert(20, "bar")
	table.Insert(30, "baz")
	size := table.Size
	if size != 3 {
		t.Fatalf("Expected the size of hash table of 3, got %v", size)
	}
}

func TestHashTableInsertOverflow(t *testing.T) {
	table := ht.CreateHashTable(3)
	table.Insert(112, "foo")
	table.Insert(222, "bar")
	table.Insert(3333, "baz")
	table.Insert(3454, "bor")
	size := table.Size
	if size != 3 {
		t.Fatalf("Expected the size of hash table of 3, got %v", size)
	}
}

func TestHashTableDelete(t *testing.T) {
	table := ht.CreateHashTable(10)
	table.Insert(111, "foo")
	table.Delete(111)
	size := table.Size
	if size != 0 {
		t.Fatalf("Expected the hash table to be empty, got thge size of %v", size)
	}

	table.Insert(111, "foo")
	table.Insert(112, "bar")
	table.Insert(122, "baz")
	table.Delete(112)
	size = table.Size
	if size != 2 {
		t.Fatalf("Expected the size of hash table of 2, got %v", size)
	}
}

func TestHashTableDeleteUnderflow(t *testing.T) {
	table := ht.CreateHashTable(10)
	result := table.Delete(111)
	if result == true {
		t.Fatalf("Expected the deletion operation to return false, got %v", result)
	}
}

func TestHashTableSearch(t *testing.T) {
	table := ht.CreateHashTable(10)
	table.Insert(123, "foo")
	table.Insert(234, "bar")
	found := table.Search(123)
	if found != "foo" {
		t.Fatalf("Expected the found value to be \"foo\", got \"%v\"", found)
	}
}
