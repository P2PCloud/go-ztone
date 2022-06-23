package one

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateKeys(t *testing.T) {
	secret, public, err := GenerateKeys()
	require.NoError(t, err)

	require.Len(t, secret, 270)
	require.Len(t, public, 141)
}
