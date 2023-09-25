package coding

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cadeusept/picture-processor/utils"
)

func MessageToCode(msg string) *[]byte {
	l := len(msg)

	runes := []rune(msg)
	code := []byte{}

	utils.Write8Bytes(rune(l), &code)

	for i := 0; i < l; i += 1 {
		utils.Write8Bytes(runes[i], &code)
	}

	return &code
}

func PutCodeIn(r *bufio.Reader, code *[]byte, w *bufio.Writer) error {
	buf_len := 1
	buf := make([]byte, buf_len)
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

	// fmt.Println(*code, len(*code), len(*code)/8, len(*code)%8)

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
