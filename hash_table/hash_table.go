package hashtable

import "math"

type Pair struct {
	k int
	v string
}

type HashTable struct {
	Capacity int
	Size     int
	list     map[int]Pair
}

func CreateHashTable(capacity int) *HashTable {
	list := make(map[int]Pair)
	return &HashTable{capacity, 0, list}
}

func (table *HashTable) String() (output string) {
	output = "[HashTable: "
	i := 0
	for _, v := range table.list {
		if i > 0 {
			output += ", "
		}
		output += v.v
		i++
	}
	output += "]"
	return output
}

func (table *HashTable) hash(key int) (index int) {
	hashConstant := 0.6180339887498949
	size := float64(table.Capacity)
	n := float64(key) * hashConstant
	index = int(math.Floor(size * n))
	return index
}

func (table *HashTable) Insert(key int, value string) {
	if table.Size == table.Capacity {
		return
	}
	index := table.hash(key)
	table.Size++
	table.list[index] = Pair{key, value}
}

func (table *HashTable) Delete(key int) {
	delete(table.list, table.hash(key))
	if table.Size == 0 {
		return
	}
	table.Size = int(math.Max(float64(table.Size)-1, 0.0))
}
