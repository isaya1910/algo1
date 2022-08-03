package hashtable

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashTable_HashFun(t *testing.T) {
	testObject := Init(17, 1)
	index := testObject.HashFun("testExample")
	fmt.Println(index)
	require.Less(t, index, 17)

	index2 := testObject.Put("testExample")
	require.Equal(t, index, index2)
}
