package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cadeusept/picture-processor/utils"
)

func main() {
	ip := "./images/"
	op := "./images/"

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
	defer ouf.Close()

	reader := bufio.NewReader(inf)
	writer := bufio.NewWriter(ouf)

	err = utils.PutCodeIn(reader, utils.MessageToCode("kto prochital tot pedik"), writer)
	if err != nil {
		fmt.Println(err)
	}
}
