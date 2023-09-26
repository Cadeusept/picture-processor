package mocks

import "fmt"

type MockWriter struct {
	Buf []byte
}

func (r *MockWriter) Write(p []byte) (n int, err error) {
	r.Buf = append(r.Buf, p...)
	fmt.Println(r.Buf)
	return len(p), nil
}
