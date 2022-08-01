package orderedlist

import (
	"constraints"
	"fmt"
)

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	_ascending bool
	count      int
}

func (l *OrderedList[T]) Count() int {
	return l.count
}

func (l *OrderedList[T]) Add(item T) {
	l.count++
	if l._ascending {
		l.addAsc(item)
		return
	}
	l.addDesc(item)
}

func (l *OrderedList[T]) addAsc(item T) {
	newNode := Node[T]{value: item}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}
	if l.head.next == nil && l.Compare(l.head.value, item) == -1 {
		l.head.next = &newNode
		newNode.prev = l.head
		l.tail = &newNode
		return
	}
	if l.head.next == nil && l.Compare(l.head.value, item) == 1 {
		tmp := l.head
		l.head = &newNode
		newNode.next = tmp
		tmp.prev = &newNode
		l.tail = tmp
		return
	}
	it := l.head
	for it != nil {
		if l.Compare(it.value, item) == 1 {
			break
		}
		it = it.next
	}
	if it == nil {
		tmp := l.tail
		l.tail = &newNode
		tmp.next = l.tail
		l.tail.prev = tmp
		return
	}

	prev := it.prev
	newNode.prev = prev
	newNode.next = it
	it.prev = &newNode
	if prev == nil {
		l.head = &newNode
		return
	}
	prev.next = &newNode
}

func (l *OrderedList[T]) addDesc(item T) {
	newNode := Node[T]{value: item}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}
	if l.head.next == nil && l.Compare(l.head.value, item) == 1 {
		l.head.next = &newNode
		newNode.prev = l.head
		l.tail = &newNode
		return
	}
	if l.head.next == nil && l.Compare(l.head.value, item) == -1 {
		tmp := l.head
		l.head = &newNode
		newNode.next = tmp
		tmp.prev = &newNode
		l.tail = tmp
		return
	}
	it := l.head
	for it != nil {
		if l.Compare(it.value, item) == -1 {
			break
		}
		it = it.next
	}
	if it == nil {
		tmp := l.tail
		l.tail = &newNode
		tmp.next = l.tail
		l.tail.prev = tmp
		return
	}

	prev := it.prev
	newNode.prev = prev
	newNode.next = it
	it.prev = &newNode
	if prev == nil {
		l.head = &newNode
		return
	}
	prev.next = &newNode
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	it := l.head
	for it != nil {
		if it.value == n {
			return *it, nil
		}
		if l._ascending && it.value > n {
			break
		}
		if !l._ascending && it.value < n {
			break
		}
	}
	return Node[T]{value: n, next: nil, prev: nil}, fmt.Errorf("item not found")
}

func (l *OrderedList[T]) Delete(n T) {
	current := l.head
	for current != nil {
		if current.value == n {
			l.count--
			prev := current.prev
			next := current.next
			if prev != nil && next != nil {
				prev.next = next
				next.prev = prev
			}
			if prev != nil && next == nil {
				prev.next = next
			}
			if prev == nil && next != nil {
				next.prev = prev
			}
			if prev == nil {
				l.head = next
			}
			if next == nil {
				l.tail = prev
			}
		}
		current = current.next
	}
}

func (l *OrderedList[T]) Clear(asc bool) {
	l._ascending = asc
	l.head = nil
	l.tail = nil
	l.count = 0
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}
