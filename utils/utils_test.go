package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddBitsTo8(t *testing.T) {
	assert.Equal(t, []string{"0", "0", "0", "0", "0", "0", "0", "0"}, AddBitsTo8([]string{}), "check again")
	assert.Equal(t, []string{"0", "0", "0", "0", "0", "0", "0", "0"}, AddBitsTo8([]string{"0"}), "check again")
	assert.Equal(t, []string{"0", "0", "0", "0", "0", "0", "0", "1"}, AddBitsTo8([]string{"1"}), "check again")
	assert.Equal(t, []string{"0", "0", "0", "0", "0", "0", "1", "0"}, AddBitsTo8([]string{"1", "0"}), "check again")
	assert.Equal(t, []string{"1", "0", "0", "0", "0", "0", "0", "0"}, AddBitsTo8([]string{"1", "0", "0", "0", "0", "0", "0", "0"}), "check again")
}

func TestWrite8Bytes(t *testing.T) {
	dest := []byte{}

	require.NoError(t, Write8Bytes('a', &dest))
	assert.Equal(t, []byte{0, 1, 1, 0, 0, 0, 0, 1}, dest)

	require.NoError(t, Write8Bytes('A', &dest))
	assert.Equal(t, []byte{
		0, 1, 1, 0, 0, 0, 0, 1,
		0, 1, 0, 0, 0, 0, 0, 1}, dest)
}

func TestChamgeTwoBits(t *testing.T) {
	src := []byte{0, 1, 1, 0, 0, 0, 0, 1} // 'a'
	var dest byte
	var i int

	require.NoError(t, ChangeTwoBits(&src, &dest, i))
	assert.Equal(t, byte(1), dest)
	i += 2

	require.NoError(t, ChangeTwoBits(&src, &dest, i))
	assert.Equal(t, byte(16), dest)
	i += 2

	require.NoError(t, ChangeTwoBits(&src, &dest, i))
	assert.Equal(t, byte(0), dest)
	i += 2

	require.NoError(t, ChangeTwoBits(&src, &dest, i))
	assert.Equal(t, byte(1), dest)
}
