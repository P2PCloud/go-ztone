package one

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateNetwork(t *testing.T) {
	c, err := NewClientFromDefaultKey()
	require.NoError(t, err)

	networks, err := c.ListNetworks()
	require.NoError(t, err)

	require.Len(t, networks, 0)
}
