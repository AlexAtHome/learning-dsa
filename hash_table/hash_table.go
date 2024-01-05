package hashtable

import (
	"fmt"
	"math"
)

type Pair struct {
	k int
	v string
}

func (p *Pair) String() string {
	return p.v
}

type HashTable struct {
	capacity int
	Size     int
	list     []*Pair
}

func CreateHashTable(capacity int) *HashTable {
	list := make([]*Pair, capacity, capacity)
	return &HashTable{capacity, 0, list}
}

func (table *HashTable) String() (output string) {
	output = "[HashTable "
	for i, v := range table.list {
		if i > 0 {
			output += ", "
		}
		if v != nil {
			output += fmt.Sprintf("%v: %v", i, v)
		} else {
			output += fmt.Sprintf("%v: ---", i)
		}
	}
	output += "]"
	return output
}

func (table *HashTable) hash(key int) int {
	hashConstant := 0.6180339887498949
	n := float64(key) * hashConstant
	return int(math.Floor(n)) % table.capacity
}

func (table *HashTable) getIndexFor(key int) (int, error) {
	index := table.hash(key)
	item := table.list[index]
	i := 1
	for item != nil && item.k != key {
		// recalculating the index, using double hashing technique
		index = (index + i*table.hash(key)) % table.capacity
		i++
		if i == table.capacity {
			return -1, fmt.Errorf("Out of space!")
		}
	}
	return index, nil
}

func (table *HashTable) Insert(key int, value string) bool {
	if table.Size == table.capacity {
		return false
	}
	index, err := table.getIndexFor(key)
	if err == nil {
		table.list[index] = &Pair{key, value}
		table.Size++
		return true
	} else {
		return false
	}
}

func (table *HashTable) Delete(key int) bool {
	if table.Size <= 0 {
		return false
	}
	index, _ := table.getIndexFor(key)
	table.list[index] = nil
	table.Size = int(math.Max(float64(table.Size-1), 0.0))
	return true
}

func (table *HashTable) Search(key int) string {
	index, _ := table.getIndexFor(key)
	found := table.list[index]
	if found == nil {
		return ""
	} else {
		return found.v
	}
}
