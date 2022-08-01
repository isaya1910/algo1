package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	testObject := LinkedList{}
	head := Node{next: nil, value: 1}
	testObject.AddInTail(head)
	assert.NotNil(t, testObject.head)
	assert.NotNil(t, testObject.tail)
	assert.Equal(t, testObject.head.value, 1)
	newNode := Node{next: nil, value: 2}
	testObject.Insert(testObject.head, newNode)
	assert.Equal(t, testObject.Count(), 2)
	node := testObject.head.next
	assert.NotNil(t, node)
	testObject.Insert(testObject.head, Node{nil, 5})
	assert.Equal(t, testObject.Count(), 3)
	assert.Equal(t, testObject.head.next.value, 5)
	assert.Equal(t, testObject.head.next.next.value, 2)
}

func TestInsertFirst(t *testing.T) {
	testObject := LinkedList{}
	testObject.InsertFirst(Node{nil, 1})
	assert.Equal(t, testObject.head.value, 1)
	testObject.InsertFirst(Node{nil, 2})
	assert.Equal(t, testObject.head.value, 2)
	assert.Equal(t, testObject.head.next.value, 1)
}

func TestDelete(t *testing.T) {
	testObject := LinkedList{}
	testObject.AddInTail(Node{nil, 1})

	testObject.AddInTail(Node{nil, 2})
	testObject.AddInTail(Node{nil, 3})

	testObject.InsertFirst(Node{nil, 2})
	testObject.InsertFirst(Node{nil, 2})
	testObject.InsertFirst(Node{nil, 2})

	assert.Equal(t, testObject.Count(), 6)
	assert.Equal(t, testObject.head.value, 2)

	testObject.Delete(2, false)
	assert.Equal(t, testObject.Count(), 5)
	testObject.Delete(2, true)
	assert.Equal(t, testObject.Count(), 2)

	testObject.Delete(3, true)
	assert.Equal(t, testObject.Count(), 1)
	assert.Equal(t, testObject.head.value, 1)
	testObject.Delete(1, true)
	assert.Equal(t, testObject.Count(), 0)
	assert.Nil(t, testObject.head)
	assert.Nil(t, testObject.tail)
}

func TestDelete2(t *testing.T) {
	testObject := LinkedList{}

	testObject.AddInTail(Node{nil, 1})
	testObject.AddInTail(Node{nil, 1})
	testObject.AddInTail(Node{nil, 1})
	testObject.AddInTail(Node{nil, 2})
	testObject.AddInTail(Node{nil, 1})
	testObject.AddInTail(Node{nil, 1})
	testObject.AddInTail(Node{nil, 1})
	testObject.AddInTail(Node{nil, 4})

	testObject.Delete(1, true)

	assert.Equal(t, testObject.Count(), 2)
	assert.Equal(t, testObject.head.value, 2)
	assert.Equal(t, testObject.tail.value, 4)

	testObject.AddInTail(Node{value: 4})

	testObject.Delete(4, false)
	assert.Equal(t, 4, testObject.tail.value)
	assert.Equal(t, 2, testObject.Count())
	testObject.Delete(4, true)
	assert.Equal(t, 2, testObject.head)
	assert.Equal(t, 2, testObject.tail)
	testObject.Delete(2, true)
	assert.Nil(t, testObject.head)
	assert.Nil(t, testObject.tail)
}

func TestFind(t *testing.T) {
	testObject := LinkedList{}
	testObject.AddInTail(Node{nil, 1})
	testObject.AddInTail(Node{nil, 5})
	actual, e := testObject.Find(1)
	assert.Nil(t, e)
	assert.Equal(t, actual.value, 1)

	actual, e = testObject.Find(2)
	assert.NotNil(t, e)
	assert.Equal(t, actual.value, -1)

	testObject.AddInTail(Node{nil, 1})
	actualNodes := testObject.FindAll(1)
	assert.Equal(t, len(actualNodes), 2)
	testObject.Clean()
	assert.Equal(t, testObject.Count(), 0)
}
