package linkedList

import (
	"errors"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

func (l *LinkedList) Count() int {
	count := 0
	iteratorNode := l.head
	for iteratorNode != nil {
		count++
		iteratorNode = iteratorNode.next
	}
	return count
}

func (l *LinkedList) Find(n int) (Node, error) {
	iteratorNode := l.head
	for iteratorNode != nil {
		if iteratorNode.value == n {
			return *iteratorNode, nil
		}
		iteratorNode = iteratorNode.next
	}
	return Node{nil, -1}, errors.New("node not found")
}

func (l *LinkedList) FindAll(n int) []Node {
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

func (l *LinkedList) Delete(n int, all bool) {
	current := l.head
	previous := l.head

	for current != nil {
		if current.value == n && current == l.head {
			current = current.next
			l.head = current
			previous = current
			if !all {
				l.UpdateTail()
				return
			}
			continue
		}
		if current.value == n {
			current = current.next
			previous.next = current
			if !all {
				l.UpdateTail()
				return
			}
			continue
		}
		if current.value != n {
			previous = current
			current = current.next
		}
	}
	l.UpdateTail()
}

func (l *LinkedList) Insert(after *Node, add Node) {
	temp := after.next
	after.next = &add
	(&add).next = temp
	l.UpdateTail()
}

func (l *LinkedList) UpdateTail() {
	iteratorNode := l.head
	l.tail = iteratorNode
	for iteratorNode != nil {
		l.tail = iteratorNode
		iteratorNode = iteratorNode.next
	}
}

func (l *LinkedList) InsertFirst(first Node) {
	temp := l.head
	l.head = &first
	(&first).next = temp
	l.UpdateTail()
}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}
