package doublelinkedlist

import (
	"errors"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) Count() int {
	count := 0
	iteratorNode := l.head
	for iteratorNode != nil {
		count++
		iteratorNode = iteratorNode.next
	}
	return count
}

func (l *LinkedList2) Find(n int) (Node, error) {
	iteratorNode := l.head
	for iteratorNode != nil {
		if iteratorNode.value == n {
			return *iteratorNode, nil
		}
		iteratorNode = iteratorNode.next
	}
	return Node{nil, nil, -1}, errors.New("node not found")
}

func (l *LinkedList2) FindAll(n int) []Node {
	var nodes []Node
	iteratorNode := l.head
	for iteratorNode != nil {
		if iteratorNode.value == n {
			nodes = append(nodes, *iteratorNode)
		}
		iteratorNode = iteratorNode.next
	}
	return nodes
}

func (l *LinkedList2) Delete(n int, all bool) {
	current := l.head
	isFind := false
	for current != nil {
		if isFind && !all {
			return
		}
		if current.value == n && current == l.head {
			current = current.next
			isFind = true
			l.head = current
			if current == nil {
				l.tail = l.head
				return
			}
			if current.next == nil {
				l.tail = current
				return
			}
			continue
		}
		if current.value == n {
			temp := current
			current = current.next
			temp.prev.next = current
			isFind = true
			if current == nil {
				l.tail = temp.prev
				return
			}
			current.prev = temp.prev
		}
		if current.next == nil {
			l.tail = current
		}
		if current.value != n {
			current = current.next
		}
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {
	temp := after.next
	after.next = &add
	add.prev = after
	(&add).next = temp
	if temp == nil {
		l.tail = &add
		return
	}
	temp.prev = &add
}

func (l *LinkedList2) InsertFirst(first Node) {
	temp := l.head
	l.head = &first
	(&first).next = temp
	if temp == nil {
		l.tail = l.head
		return
	}
	temp.prev = &first
	if temp.next == nil {
		l.tail = temp
	}
}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}
