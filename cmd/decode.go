package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var decodeJSON bool
var decodeOutput string

var decode = &cobra.Command{
	Use:   "decode [jwt]",
	Short: "Decode a JWT",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Decoding...")
	},
}

func init() {
	decode.Flags().BoolVar(&decodeJSON, "json", false, "Output as json")
	decode.Flags().StringVarP(&decodeOutput, "output", "o", "", "Go template string to format the output")
	root.AddCommand(decode)
}
