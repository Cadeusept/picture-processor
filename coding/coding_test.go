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

func TestPutCodeIn(t *testing.T) {
	//r := bufio.NewReader(rd io.Reader)
}
