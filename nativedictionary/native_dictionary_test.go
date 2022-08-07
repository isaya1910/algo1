package nativedictionary

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNativeDictionary(t *testing.T) {
	testObject := Init[int](10)

	require.False(t, testObject.IsKey("test"))
	testObject.Put("test", 1)
	require.True(t, testObject.IsKey("test"))
	actual, err := testObject.Get("test")
	require.NoError(t, err)
	require.Equal(t, 1, actual)

	actual, err = testObject.Get("test1")
	require.Error(t, err)
	require.Empty(t, actual)
}
