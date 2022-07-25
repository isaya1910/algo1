package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack_Peek(t *testing.T) {
	testObject := Stack[int]{}
	actual, err := testObject.Peek()
	require.Error(t, err)
	require.Empty(t, actual)
	testObject.Push(8)
	actual, err = testObject.Peek()
	require.NoError(t, err)
	require.Equal(t, 8, actual)
}

func TestStack_Pop(t *testing.T) {
	testObject := Stack[int]{}
	actual, err := testObject.Pop()
	require.Empty(t, actual)
	require.Error(t, err)
	testObject.Push(8)
	actual, err = testObject.Pop()
	require.Equal(t, 8, actual)
	require.Equal(t, 0, testObject.Size())
	require.NoError(t, err)
}

func TestStack_Size(t *testing.T) {
	testObject := Stack[int]{}
	for i := 0; i < 32; i++ {
		testObject.Push(i)
	}
	require.Equal(t, 32, testObject.Size())
	for i := 0; i < 32; i++ {
		actual, err := testObject.Pop()
		require.NoError(t, err)
		require.Equal(t, 31-i, actual)
		require.Equal(t, 31-i, testObject.Size())
	}
}
