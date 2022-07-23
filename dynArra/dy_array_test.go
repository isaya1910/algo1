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

	require.Equal(t, 30, testDynArray.count)

	testDynArray.Insert(10, 30)
	actual, err := testDynArray.GetItem(30)
	require.NoError(t, err)
	require.Equal(t, 10, actual)
	require.Equal(t, 31, testDynArray.count)
}

func TestDynArrayCapacity(t *testing.T) {
	var testDynArray DynArray[int]
	testDynArray.Init()
	require.Equal(t, 16, testDynArray.capacity)
	for i := 0; i < 17; i++ {
		testDynArray.Insert(1, testDynArray.count)
	}
	require.Equal(t, 17, testDynArray.count)
	require.Equal(t, 32, testDynArray.capacity)

	err := testDynArray.Insert(48, testDynArray.count)
	require.NoError(t, err)

	actual, err := testDynArray.GetItem(testDynArray.count - 1)
	require.NoError(t, err)
	require.Equal(t, 48, actual)
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
