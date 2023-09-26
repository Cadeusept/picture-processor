package mocks

import "io"

type MockReader struct {
	Buf []byte
}

func (b *MockReader) Read(p []byte) (n int, err error) {
	n = copy(p, b.Buf)
	return n, io.EOF
}
