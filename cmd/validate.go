package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var validateIss string
var validateSub string
var validateAud string
var validateExp string
var validateNbf string
var validateIat string
var validateJti string
var validateAlgorithm string
var validateSecret string

var validate = &cobra.Command{
	Use:   "validate [jwt]",
	Short: "Validate a JWT",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Validating...")
	},
}

func init() {
	validate.Flags().StringVar(&validateIss, "iss", "", "Fails if Issuer claim does not match")
	validate.Flags().StringVar(&validateSub, "sub", "", "Fails if Subject claim does not match")
	validate.Flags().StringVar(&validateAud, "aud", "", "Fails if Audience claim does not match")
	validate.Flags().StringVar(&validateExp, "exp", "", "Fails if Expiration Time claim is before this value")
	validate.Flags().StringVar(&validateNbf, "nbf", "", "Fails of Not Before claim is after this value")
	validate.Flags().StringVar(&validateIat, "iat", "", "Fails if Issued At claim is after this value")
	validate.Flags().StringVar(&validateJti, "jti", "", "Fails if JWT ID claim does not match")
	validate.Flags().StringVarP(&validateAlgorithm, "algorithm", "a", "", "The algorithm to validate against. Fails on mismatch")
	validate.Flags().StringVarP(&validateSecret, "secret", "s", "", "The secret (string, @file, or '-' to read from stdin). Fails if signature is invalid")
	root.AddCommand(validate)
}
