package cmd

import (
	"fmt"
	"os"

	"github.com/compiledpanda/jwt/internal"
	"github.com/spf13/cobra"
)

var decodeOptions internal.DecodeOptions

var decode = &cobra.Command{
	Use:   "decode [jwt]",
	Short: "Decode a JWT",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		str, err := internal.Decode(args[0], decodeOptions)
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
