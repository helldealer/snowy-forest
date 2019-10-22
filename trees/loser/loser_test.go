package loser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type elem int

func (e elem) Lose(LoserElem Elem) bool {
	return e >= LoserElem.(elem)
}

var initLoserElem = []Elem{elem(-2), elem(6), elem(3), elem(-1), elem(1), elem(5), elem(4)}

func TestLoser(t *testing.T) {
	loser := New(initLoserElem, elem(-10))
	require.Equal(t, loser.r[0], 0)
	require.Equal(t, loser.p[loser.r[0]], elem(-2))

	index, e := loser.Winner()
	require.Equal(t, index, 0)
	require.Equal(t, e, elem(-2))

	index, e = loser.Update(index, elem(0))
	require.Equal(t, index, 3)
	require.Equal(t, e, elem(-1))

	res := loser.Iterate(elem(100))
	require.Equal(t, len(res), 7)
	require.Equal(t, res[0], elem(-1))
	require.Equal(t, res[1], elem(0))
	require.Equal(t, res[2], elem(1))
	require.Equal(t, res[3], elem(3))
	require.Equal(t, res[4], elem(4))
	require.Equal(t, res[5], elem(5))
	require.Equal(t, res[6], elem(6))
}
