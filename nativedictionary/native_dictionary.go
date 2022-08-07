package nativedictionary

import (
	"fmt"
	"hash/fnv"
)

type NativeDictionary[T any] struct {
	size   int
	slots  []string
	values []T
}

func Init[T any](sz int) NativeDictionary[T] {
	nd := NativeDictionary[T]{size: sz, slots: nil, values: nil}
	nd.slots = make([]string, sz)
	nd.values = make([]T, sz)
	return nd
}

func (nd *NativeDictionary[T]) HashFun(value string) int {
	h := fnv.New32a()
	h.Write([]byte(value))
	return int(h.Sum32()) % nd.size
}

func (nd *NativeDictionary[T]) IsKey(key string) bool {
	for i := 0; i < nd.size; i++ {
		if nd.slots[i] == key {
			return true
		}
	}
	return false
}

func (nd *NativeDictionary[T]) Get(key string) (T, error) {
	var result T
	index := nd.HashFun(key)
	if nd.slots[index] == "" {
		return result, fmt.Errorf("item not found")
	}
	result = nd.values[index]
	return result, nil
}

func (nd *NativeDictionary[T]) Put(key string, value T) {
	index := nd.HashFun(key)
	nd.slots[index] = key
	nd.values[index] = value
}
