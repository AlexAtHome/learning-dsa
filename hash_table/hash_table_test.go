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
	if table.Size != 3 {
		t.Fatalf("Expected the size of hash table of 3, got %v", table.Size)
	}
}

func TestHashTableDelete(t *testing.T) {
	table := ht.CreateHashTable(10)
	table.Insert(111, "foo")
	table.Delete(111)
	if table.Size != 0 {
		t.Fatalf("Expected the hash table to be empty, got thge size of %v", table.Size)
	}

	table.Insert(111, "foo")
	table.Insert(112, "bar")
	table.Insert(122, "baz")
	table.Delete(112)
	if table.Size != 2 {
		t.Fatalf("Expected the size of hash table of 2, got %v", table.Size)
	}
}
