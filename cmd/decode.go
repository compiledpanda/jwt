package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/compiledpanda/jwt/internal"
	"github.com/spf13/cobra"
)

var decodeOptions internal.DecodeOptions

var decode = &cobra.Command{
	Use:   "decode [jwt]",
	Short: "Decode a JWT",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// JWT
		var jwt []byte
		if args[0] == "-" {
			jwt = readStdIn()
		} else if strings.HasPrefix(args[0], "@") {
			jwt = readFile(args[0][1:])
		} else {
			jwt = []byte(args[0])
		}
		jwt = bytes.TrimSpace(jwt)

		str, err := internal.Decode(jwt, decodeOptions)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(str)
	},
}

func init() {
	decode.Flags().BoolVar(&decodeOptions.JSON, "json", false, "Output as json")
	decode.Flags().StringVarP(&decodeOptions.Output, "output", "o", "", "Go template string to format the output")
	root.AddCommand(decode)
}
