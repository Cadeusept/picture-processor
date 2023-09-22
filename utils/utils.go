package utils

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func changeLastBit(src byte, dest *byte) error {
	var res byte
	dest_str := strings.Split(strconv.FormatInt(int64(*dest), 2), "")
	dest_str[len(dest_str)-1] = fmt.Sprintf("%v", src)
	for i := 0; i < len(dest_str); i++ {
		bit, err := strconv.Atoi(dest_str[i])
		if err != nil {
			return err
		}

		if bit == 1 {
			res += byte(math.Pow(2, float64(i)))
		}
	}
	*dest = res
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

	for {
		_, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if i < cl {
			err := changeLastBit((*code)[i], &buf[buf_len-1])
			if err != nil {
				return fmt.Errorf("error changing last bit: %w", err)
			}
		}

		_, err = w.Write(buf)
		if err != nil {
			return err
		}
		i++
	}

	return nil
}
