package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cadeusept/picture-processor/coding"
	"github.com/cadeusept/picture-processor/decoding"
)

func main() {
	ip := "./../images/"
	op := "./../images/"

	inf, err := os.Open(ip + "parrot.bmp")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer inf.Close()

	ouf, err := os.Create(op + "myparrot.bmp")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}

	reader := bufio.NewReader(inf)
	writer := bufio.NewWriter(ouf)

	err = coding.PutCodeIn(reader, coding.MessageToCode("kto pro4el beast"), writer)
	if err != nil {
		fmt.Println(err)
		return
	}

	ouf.Close()

	ouf, err = os.Open(op + "myparrot.bmp")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer ouf.Close()

	reader = bufio.NewReader(ouf)

	msg, err := decoding.PutCodeOut(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("coded message: %v\n", msg)
}
