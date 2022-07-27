package queue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	testObject := Queue[int]{}

	require.Equal(t, 0, testObject.Size())
	testObject.Enqueue(1)
	require.Equal(t, 1, testObject.Size())

	actual, err := testObject.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 1, actual)
	require.Equal(t, testObject.Size(), 0)

	for i := 0; i < 30; i++ {
		testObject.Enqueue(i)
	}
	require.Equal(t, 30, testObject.Size())
	for i := 0; i < 30; i++ {
		actual, err = testObject.Dequeue()
		require.Equal(t, i, actual)
		require.Equal(t, 29-i, testObject.Size())
	}

}
