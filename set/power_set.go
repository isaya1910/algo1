package set

import (
	"constraints"
)

type PowerSet[T constraints.Ordered] struct {
	dict map[T]T
}

func Init[T constraints.Ordered](sz int) PowerSet[T] {
	powerSet := PowerSet[T]{}
	powerSet.dict = make(map[T]T)
	return powerSet
}

func (p *PowerSet[T]) Size() int {
	return len(p.dict)
}

func (p *PowerSet[T]) Put(value T) {
	p.dict[value] = value
}

func (p *PowerSet[T]) Get(value T) bool {
	if _, isExist := p.dict[value]; isExist {
		return true
	}
	return false
}

func (p *PowerSet[T]) Remove(value T) bool {
	if _, isExist := p.dict[value]; isExist {
		delete(p.dict, value)
		return true
	}
	return false
}

func (p *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	// пересечение текущего множества и set2
	var result PowerSet[T]
	result.dict = make(map[T]T)
	for _, key := range p.dict {
		if _, isExist := set2.dict[key]; isExist {
			result.Put(key)
		}
	}
	return result
}

func (p *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	// объединение текущего множества и set2
	var result PowerSet[T]
	result.dict = make(map[T]T)
	for _, key := range p.dict {
		result.Put(key)
	}
	for _, key := range set2.dict {
		result.Put(key)
	}
	return result
}

func (p *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T]
	result.dict = make(map[T]T)

	for _, key := range p.dict {
		if _, isExist := set2.dict[key]; !isExist {
			result.Put(key)
		}
	}

	return result
}

func (p *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
	for _, key := range set2.dict {
		if _, isExist := p.dict[key]; !isExist {
			return false
		}
	}
	return true
}
