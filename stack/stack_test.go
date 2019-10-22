package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack(t *testing.T) {
	s := New(100)
	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	require.Equal(t, 100, s.Len())
	for i := 99; i >= 0; i-- {
		require.Equal(t, i, s.Pop())
	}
	require.Equal(t, 0, s.Len())
}