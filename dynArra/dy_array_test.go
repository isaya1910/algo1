package dynArra

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDynArray_Remove(t *testing.T) {
	var testDynArray DynArray[int]
	testDynArray.Init()

	for i := 0; i < 33; i++ {
		testDynArray.Insert(i, 0)
	}

	require.Equal(t, 64, testDynArray.capacity)

	testDynArray.Remove(3)
	testDynArray.Remove(5)
	testDynArray.Remove(6)

	require.Equal(t, 42, testDynArray.capacity)

}

func TestDynArray_Insert(t *testing.T) {
	var testDynArray DynArray[int]
	testDynArray.Init()

	for i := 0; i < 32; i++ {
		testDynArray.Append(i)
	}

	actual, err := testDynArray.GetItem(0)
	require.NoError(t, err)

	require.Equal(t, 0, actual)

	require.Equal(t, 32, testDynArray.capacity)

	err = testDynArray.Insert(44, 3)
	require.NoError(t, err)
	actual, err = testDynArray.GetItem(3)
	require.NoError(t, err)
	require.Equal(t, 44, actual)

	require.Equal(t, 64, testDynArray.capacity)

}
