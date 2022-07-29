package deque

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeque_(t *testing.T) {
	testObject := Deque[int]{}
	testObject.AddFront(1)
	require.Equal(t, 1, testObject.Size())
	testObject.AddTail(9)
	actual, err := testObject.RemoveFront()
	require.NoError(t, err)
	require.Equal(t, 1, actual)
	testObject.AddTail(10)
	actual, err = testObject.RemoveFront()
	require.NoError(t, err)
	require.Equal(t, 9, actual)
	testObject.AddFront(100)

	actual, err = testObject.RemoveTail()
	require.NoError(t, err)
	require.Equal(t, 10, actual)
}
