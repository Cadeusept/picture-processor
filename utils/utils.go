package utils

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func writing8Bytes(symbol rune, dest *[]byte) error {
	src := strings.Split(strconv.FormatInt(int64(symbol), 2), "")
	if len(src) < 8 {
		res := []string{}
		for i := 0; i < 8-len(src); i++ {
			res = append(res, "0")
		}
		res = append(res, src...)
		src = res
	}

	for _, v := range src {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*dest = append(*dest, byte(i))
	}

	return nil
}

func MessageToCode(msg string) *[]byte {
	l := len(msg)

	runes := []rune(msg)
	code := []byte{}

	writing8Bytes(rune(l), &code)

	for i := 0; i < l; i += 1 {
		writing8Bytes(runes[i], &code)
	}

	return &code
}

func PutCodeIn(r *bufio.Reader, code *[]byte, w *bufio.Writer) error {
	buf := make([]byte, 8)
	i := 0
	cl := len(*code)
	for {
		_, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if i < cl {
			buf[7] = (*code)[i]
			i++
		}

		_, err = w.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}
