package hashtable

import "hash/fnv"

type HashTable struct {
	size  int
	step  int
	slots []string
}

func Init(sz int, stp int) HashTable {
	ht := HashTable{size: sz, step: stp, slots: nil}
	ht.slots = make([]string, sz)
	return ht
}

func (ht *HashTable) HashFun(value string) int {
	h := fnv.New32a()
	h.Write([]byte(value))
	return int(h.Sum32()) % ht.size
}

func (ht *HashTable) SeekSlot(value string) int {
	index := ht.HashFun(value)
	for index < ht.size {
		if ht.slots[index] == "" || ht.slots[index] == value {
			return index
		}
		index += ht.step
	}
	return -1
}

func (ht *HashTable) Put(value string) int {
	index := ht.SeekSlot(value)
	if index == -1 {
		return -1
	}
	ht.slots[index] = value
	return index
}

func (ht *HashTable) Find(value string) int {
	index := ht.SeekSlot(value)
	if index == -1 {
		return -1
	}
	if ht.slots[index] != value {
		return -1
	}
	return index
}
