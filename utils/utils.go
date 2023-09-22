package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Write8Bytes(symbol rune, dest *[]byte) error {
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

func ChangeTwoBits(src *[]byte, dest *byte, i int) (int, error) {
	var res byte
	var bitsNum int
	dest_str := strings.Split(strconv.FormatInt(int64(*dest), 2), "")
	for len(dest_str) != 8 {
		dest_str = append([]string{"0"}, dest_str...)
	}

	// if len(*src) > i+2 {
	dest_str[(len(dest_str)/2)-1] = fmt.Sprintf("%v", (*src)[i])
	dest_str[len(dest_str)-1] = fmt.Sprintf("%v", (*src)[i+1])
	bitsNum = 2
	// } else {
	// 	dest_str[(len(dest_str)/2)-1] = fmt.Sprintf("%v", (*src)[i])
	// 	bitsNum = 1
	// }

	for i := 0; i < len(dest_str); i++ {
		bit, err := strconv.Atoi(dest_str[i])
		if err != nil {
			return 0, err
		}

		if bit == 1 {
			res += byte(math.Pow(2, float64(7-i%8)))
		}
	}

	*dest = res
	return bitsNum, nil
}
