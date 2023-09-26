package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add missing zeroes up to 8 elements
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

// Append symbol in bit representation to dest
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

// Replaces middle and end bits in dest with values from src on positions i and i+1
func ChangeTwoBits(src *[]byte, dest *byte, i int) error {
	var res byte
	dest_str := AddBitsTo8(strings.Split(strconv.FormatInt(int64(*dest), 2), ""))

	dest_str[(len(dest_str)/2)-1] = fmt.Sprintf("%v", (*src)[i])
	dest_str[len(dest_str)-1] = fmt.Sprintf("%v", (*src)[i+1])

	for i := 0; i < len(dest_str); i++ {
		bit, err := strconv.Atoi(dest_str[i])
		if err != nil {
			return err
		}

		if bit == 1 {
			res += byte(math.Pow(2, float64(7-i%8)))
		}
	}

	*dest = res
	return nil
}
