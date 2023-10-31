package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// AddBitsTo8 adds missing zeroes up to 8 elements
func AddBitsTo8(str []string) []string {
	if len(str) < 8 {
		res := []string{}
		for i := 0; i < 8-len(str); i++ {
			res = append(res, "0")
		}
		res = append(res, str...)
		str = res
	}
	return str
}

// Write8Bits appends symbol in bit representation to dest
func Write8Bits(symbol rune, dest *[]byte) error {
	src := AddBitsTo8(strings.Split(strconv.FormatInt(int64(symbol), 2), ""))

	for _, v := range src {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*dest = append(*dest, byte(i))
	}

	return nil
}

// ChangeTwoBits replaces middle and end bits in dest with values from src on positions i and i+1
func ChangeTwoBits(src *[]byte, dest *byte, i int) error {
	var res byte
	destStr := AddBitsTo8(strings.Split(strconv.FormatInt(int64(*dest), 2), ""))

	destStr[(len(destStr)/2)-1] = fmt.Sprintf("%v", (*src)[i])
	destStr[len(destStr)-1] = fmt.Sprintf("%v", (*src)[i+1])

	for i := 0; i < len(destStr); i++ {
		bit, err := strconv.Atoi(destStr[i])
		if err != nil {
			return err
		}

		if bit == 1 {
			res += 1 << (7 - i%8)
		}
	}

	*dest = res
	return nil
}
