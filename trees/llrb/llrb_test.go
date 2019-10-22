package llrb

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := New()
	for i := 0; i < 1000; i++ {
		tree.Insert(elem(i))
	}

	for i := 2000; i > 0; i-- {
		tree.Insert(elem(i))
	}
	sum := 0
	tree.Iterate(nil, nil, true, func(e Elem) bool {
		sum++
		return true
	})
	require.Equal(t, 2001, sum)
	require.Equal(t, elem(2000), tree.Max())
	require.Equal(t, elem(0), tree.Min())
	require.Equal(t, 2001, tree.Len())
	require.Equal(t, elem(100), tree.Get(elem(100)))
	require.True(t, tree.Has(elem(10)))
	require.False(t,tree.Update(elem(3000)))

	tree.InsertOrUpdate(elem(3000))
	require.Equal(t, elem(3000), tree.Get(elem(3000)))
	tree.InsertOrUpdate(elem(100))
	require.Equal(t, elem(100), tree.Get(elem(100)))
}

type elem int

func (e elem) Compare(elem2 Elem) int {
	if e < elem2.(elem) {
		return -1
	} else if e == elem2.(Elem) {
		return 0
	}
	return 1
}
