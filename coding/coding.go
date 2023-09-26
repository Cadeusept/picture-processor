package coding

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cadeusept/picture-processor/utils"
)

// MessageToCode generates code slice - bit representation of msg and its size
func MessageToCode(msg string) *[]byte {
	l := len(msg)

	runes := []rune(msg)
	code := []byte{}

	utils.Write8Bits(rune(l), &code)

	for i := 0; i < l; i++ {
		utils.Write8Bits(runes[i], &code)
	}

	return &code
}

// InsertBmpCodeIn inserts code inside .bmp picture
func InsertBmpCodeIn(r *bufio.Reader, code *[]byte, w *bufio.Writer) error {
	bufLen := 1
	buf := make([]byte, bufLen)
	i := 0
	cl := len(*code)
	if true {
		garbageBuf := make([]byte, 54)
		_, err := r.Read(garbageBuf)
		if err != nil {
			return err
		}

		_, err = w.Write(garbageBuf)
		if err != nil {
			return err
		}
	}

	for {
		_, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if i < cl {
			if utils.ChangeTwoBits(code, &buf[0], i) != nil {
				return fmt.Errorf("error changing last bit: %w", err)
			}
			i += 2
		}

		_, err = w.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}
