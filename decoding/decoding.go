package decoding

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/cadeusept/picture-processor/utils"
)

func earnBits(src byte) (*[]byte, error) {
	src_str := utils.AddBitsTo8(strings.Split(strconv.FormatInt(int64(src), 2), ""))

	bit1, err := strconv.Atoi(src_str[len(src_str)/2-1])
	if err != nil {
		return nil, err
	}

	bit2, err := strconv.Atoi(src_str[len(src_str)-1])
	if err != nil {
		return nil, err
	}

	return &[]byte{byte(bit1), byte(bit2)}, nil
}

func PutCodeOut(r *bufio.Reader) (string, error) {
	code := []byte{}
	buf := make([]byte, 1)
	nSlice := []byte{}
	var n int

	if true {
		garbageBuf := make([]byte, 54)
		_, err := r.Read(garbageBuf)
		if err != nil {
			return "", err
		}
	}

	for i := 0; i < 4; i++ {
		_, err := r.Read(buf)
		if err != nil {
			return "", err
		}

		tmpSlice, err := earnBits(buf[0])
		if err != nil {
			return "", err
		}

		nSlice = append(nSlice, *tmpSlice...)
	}

	for i, v := range nSlice {
		if v == 1 {
			n += int(math.Pow(2, float64(len(nSlice)-i-1)))
		}
	}

	for i := 0; i < n*8; i += 2 {
		_, err := r.Read(buf)
		if err != nil {
			return "", err
		}

		slc, err := earnBits(buf[0])
		if err != nil {
			return "", err
		}

		code = append(code, (*slc)...)

	}

	var res string
	var tmpRes int
	j := 0
	for i, v := range code {
		if v == 1 {
			tmpRes += int(math.Pow(2, float64(7-(i%8))))
		}
		j++
		if j == 8 {
			j = 0
			res += string(rune(tmpRes))
			tmpRes = 0
		}
	}

	return string(res), nil
}
