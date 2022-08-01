package orderedlist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderedList_Add_Asc(t *testing.T) {
	testObject := OrderedList[int]{}
	testObject._ascending = true
	require.Equal(t, 0, testObject.count)
	testObject.Add(1)
	require.Equal(t, 1, testObject.count)
	require.Equal(t, 1, testObject.tail.value)
	require.Equal(t, 1, testObject.head.value)

	testObject.Add(5)
	require.Equal(t, 1, testObject.head.value)
	require.Equal(t, 5, testObject.tail.value)

	testObject.Add(7)
	require.Equal(t, 1, testObject.head.value)
	require.Equal(t, 7, testObject.tail.value)
	require.Equal(t, 5, testObject.head.next.value)
	require.Equal(t, 5, testObject.tail.prev.value)
}

func TestOrderedList_Add_Des(t *testing.T) {
	testObject := OrderedList[int]{}
	require.Equal(t, 0, testObject.count)
	testObject.Add(1)
	require.Equal(t, 1, testObject.count)
	require.Equal(t, 1, testObject.tail.value)
	require.Equal(t, 1, testObject.head.value)

	testObject.Add(5)
	require.Equal(t, 5, testObject.head.value)
	require.Equal(t, 1, testObject.tail.value)

	testObject.Add(7)
	require.Equal(t, 7, testObject.head.value)
	require.Equal(t, 1, testObject.tail.value)
	require.Equal(t, 5, testObject.head.next.value)
	require.Equal(t, 5, testObject.tail.prev.value)

	actual, err := testObject.Find(-3)
	require.Error(t, err)
	actual, err = testObject.Find(7)
	require.Equal(t, 7, actual.value)
	require.NoError(t, err)
}
