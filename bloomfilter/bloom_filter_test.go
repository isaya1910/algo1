package bloomfilter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBloomFilter_Add(t *testing.T) {
	filterLen := 32
	testObject := BloomFilter{
		filter_len: filterLen,
	}
	require.False(t, testObject.IsValue("0123456789"))
	testObject.Add("0123456789")
	require.True(t, testObject.IsValue("0123456789"))
}
