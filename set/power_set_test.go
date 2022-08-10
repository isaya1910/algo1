package set

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPowerSet_Put(t *testing.T) {
	testObject := Init[int](10)
	testObject.Put(1)
	require.Equal(t, 1, testObject.Size())
	require.True(t, testObject.Get(1))
	require.False(t, testObject.Get(2))
	testObject.Put(2)
	require.True(t, testObject.Get(2))
	require.Equal(t, 2, testObject.Size())

	require.True(t, testObject.Remove(2))
	require.Equal(t, 1, testObject.Size())
	require.False(t, testObject.Remove(2))
}

func TestPowerSet_Intersection(t *testing.T) {
	testObject := Init[int](10)

	testObject.Put(1)
	testObject.Put(2)
	testObject.Put(3)

	set := Init[int](10)
	set.Put(2)
	set.Put(3)
	set.Put(4)

	actual := testObject.Intersection(set)

	require.Equal(t, 2, actual.Size())

	actual = testObject.Union(set)

	require.Equal(t, 4, actual.Size())

	actual = testObject.Difference(set)

	require.Equal(t, 1, actual.Size())
	require.True(t, actual.Get(1))

	require.False(t, testObject.IsSubset(set))

	set2 := Init[int](10)
	set2.Put(1)
	set2.Put(2)

	require.True(t, testObject.IsSubset(set2))

	set2.Put(5)
	require.False(t, testObject.IsSubset(set2))
	set2.Remove(5)

	require.True(t, testObject.IsSubset(set2))
}
