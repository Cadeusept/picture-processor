package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cadeusept/picture-processor/coding"
	"github.com/cadeusept/picture-processor/decoding"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ip string
var op string
var flagDecode bool
var flagCode bool

var rootCmd = cobra.Command{
	Use:     "picture-processor",
	Version: "v1.0.0",
	Short:   "It's picture processor",
	Long:    "You can process the pictures",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if flagCode || !flagDecode {
			if len(args) == 0 {
				fmt.Println(fmt.Errorf("error: please, provide a message"))
				return
			}

			if filepath.Ext(ip) != filepath.Ext(op) {
				fmt.Println(fmt.Errorf("error: input and output extentions should be the same"))
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

			ext := filepath.Ext(ip)
			switch ext {
			case ".bmp":
				err = coding.InsertBmpCodeIn(reader, coding.MessageToCode(args[0]), writer)
				if err != nil {
					fmt.Println(err)
					return
				}
			default:
				fmt.Println(fmt.Errorf("error: extension %s not supported", ext))
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

			var msg string
			ext := filepath.Ext(ip)
			switch ext {
			case ".bmp":
				msg, err = decoding.GetBmpMsgOut(reader)
				if err != nil {
					fmt.Println(err)
					return
				}
			default:
				fmt.Println(fmt.Errorf("error: extension %s not supported", ext))
				return
			}

			fmt.Printf("coded message: %v\n", msg)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&ip, "source", "s", "./images/parrot.bmp", "Change input source path on provided")
	rootCmd.Flags().StringVarP(&op, "output", "o", "./images/myparrot.bmp", "Change coded file output path on provided")
	rootCmd.Flags().BoolVarP(&flagCode, "code", "c", false, "Code file")
	rootCmd.Flags().BoolVarP(&flagDecode, "decode", "d", false, "Decode file")
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error executing command: %v", err.Error())
	}
}
