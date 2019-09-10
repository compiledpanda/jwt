package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "jwt",
	Short: "Encode, Decode, and Validate JWTs",
}

// Execute the command
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
