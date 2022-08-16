package nativecache

import (
	"fmt"
	"hash/fnv"
)

type NativeCache[T any] struct {
	size   int
	slots  []string
	values []T
	hits   []int
}

func Init[T any](sz int) NativeCache[T] {
	nc := NativeCache[T]{size: sz, slots: nil, values: nil}
	nc.slots = make([]string, sz)
	nc.values = make([]T, sz)
	nc.hits = make([]int, sz)
	return nc
}

func (nc *NativeCache[T]) HashFun(value string) int {
	h := fnv.New32a()
	h.Write([]byte(value))
	return int(h.Sum32()) % nc.size
}

func (nc *NativeCache[T]) SeekSlot(value string) int {
	index := nc.HashFun(value)

	minHits := nc.hits[0]
	minHitsIndex := 0
	for index < nc.size {
		if nc.slots[index] == "" || nc.slots[index] == value {
			return index
		}
		if nc.hits[index] < minHits {
			minHits = nc.hits[index]
			minHitsIndex = index
		}
		index++
	}
	return minHitsIndex
}

func (nc *NativeCache[T]) IsKey(key string) bool {
	for i := 0; i < nc.size; i++ {
		if nc.slots[i] == key {
			return true
		}
	}
	return false
}

func (nc *NativeCache[T]) Get(key string) (T, error) {
	var result T
	index := nc.HashFun(key)
	if !nc.IsKey(key) {
		return result, fmt.Errorf("item not found")
	}
	nc.hits[index]++
	result = nc.values[index]
	return result, nil
}

func (nc *NativeCache[T]) Put(key string, value T) {
	index := nc.SeekSlot(key)
	nc.slots[index] = key
	nc.values[index] = value
}
