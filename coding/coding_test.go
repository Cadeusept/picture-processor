package coding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageToCode(t *testing.T) {
	assert.Equal(t, []byte{
		0, 0, 0, 0, 0, 0, 0, 1, // n
		0, 1, 1, 0, 0, 0, 0, 1}, *MessageToCode("a"))

	assert.Equal(t, []byte{
		0, 0, 0, 0, 0, 0, 1, 0, // n
		0, 1, 0, 0, 0, 0, 0, 1,
		0, 1, 1, 0, 0, 0, 0, 1}, *MessageToCode("Aa"))
}

/* TODO
func TestPutCodeIn(t *testing.T) {
	buf := []byte{0, 10, 20, 26}
	code := []byte{1, 2}
	mr := mocks.MockReader{Buf: buf}
	mw := mocks.MockWriter{Buf: []byte{}}
	r := bufio.NewReader(&mr)
	w := bufio.NewWriter(&mw)

	require.NoError(t, PutCodeIn(r, &code, w))

	assert.Equal(t, []byte{1, 26, 20}, mw.Buf)
}
*/
