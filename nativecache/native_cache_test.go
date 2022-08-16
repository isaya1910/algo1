package nativecache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNativeCache(t *testing.T) {
	testObject := Init[int](10)
	testObject.Put("test", 1)
	actual, err := testObject.Get("test")
	require.NoError(t, err)
	require.Equal(t, 1, actual)
	index := testObject.HashFun("test")
	testObject.Get("test")
	require.Equal(t, testObject.hits[index], 2)
}
