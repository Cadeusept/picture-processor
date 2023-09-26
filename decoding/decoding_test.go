package decoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEarnBits(t *testing.T) {
	var arr *[]byte
	var err error

	arr, err = earnBits(128)
	require.NoError(t, err)
	assert.Equal(t, []byte{0, 0}, *arr)

	arr, err = earnBits(16)
	require.NoError(t, err)
	assert.Equal(t, []byte{1, 0}, *arr)

	arr, err = earnBits(33)
	require.NoError(t, err)
	assert.Equal(t, []byte{0, 1}, *arr)

	arr, err = earnBits(17)
	require.NoError(t, err)
	assert.Equal(t, []byte{1, 1}, *arr)
}
