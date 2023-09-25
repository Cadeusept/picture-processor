package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cadeusept/picture-processor/coding"
	"github.com/cadeusept/picture-processor/decoding"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ip string
var op string
var flagDecode bool

var rootCmd = cobra.Command{
	Use:     "picture-processor",
	Version: "v1.0.0",
	Short:   "It's picture processor",
	Long:    "You can process the pictures",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if !flagDecode {
			if len(args) == 0 {
				fmt.Println(fmt.Errorf("error: please, provide a message"))
				return
			}

			inf, err := os.Open(ip)
			if err != nil {
				fmt.Println("Unable to open file:", err)
				return
			}
			defer inf.Close()

			ouf, err := os.Create(op)
			if err != nil {
				fmt.Println("Unable to open file:", err)
				return
			}
			defer ouf.Close()

			reader := bufio.NewReader(inf)
			writer := bufio.NewWriter(ouf)

			err = coding.PutCodeIn(reader, coding.MessageToCode(args[0]), writer)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("message coded successfully into file: %s\n", op)
		} else {
			ipf, err := os.Open(ip)
			if err != nil {
				fmt.Println("Unable to open file:", err)
				return
			}
			defer ipf.Close()

			reader := bufio.NewReader(ipf)

			msg, err := decoding.PutCodeOut(reader)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("coded message: %v\n", msg)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&ip, "ipath", "i", "./../images/parrot.bmp", "Change input path on provided")
	rootCmd.Flags().StringVarP(&op, "opath", "o", "./../images/myparrot.bmp", "Change output path on provided")
	rootCmd.Flags().BoolVarP(&flagDecode, "decode", "d", false, "Change prog mode")
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error executing command: %v", err.Error())
	}
}
