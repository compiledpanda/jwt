package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/compiledpanda/jwt/internal"
	"github.com/spf13/cobra"
)

var validateIss string
var validateSub string
var validateAud []string
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
		options := internal.ValidateOptions{}

		// Validators
		if validateIss != "" {
			options.Issuer = validateIss
		}
		if validateSub != "" {
			options.Subject = validateSub
		}
		if len(validateAud) > 0 {
			options.Audience = validateAud
		}
		if validateExp != "" {
			// TODO
		}
		if validateNbf != "" {
			// TODO
		}
		if validateIat != "" {
			// TODO
		}
		if validateJti != "" {
			options.Jti = validateJti
		}

		// Algorithm
		if !internal.ValidAlgorithm(validateAlgorithm) {
			fmt.Println("algorithm is invalid. Must be one of ", strings.Join(internal.Algorithms, ", "))
			os.Exit(1)
		}
		options.Algorithm = validateAlgorithm

		// Secret
		if validateSecret == "-" {
			options.Secret = readStdIn()
		} else if strings.HasPrefix(validateSecret, "@") {
			options.Secret = readFile(validateSecret[1:])
		} else {
			options.Secret = []byte(validateSecret)
		}

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

		str, err := internal.Validate(jwt, options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(str)
	},
}

func init() {
	validate.Flags().StringVar(&validateIss, "iss", "", "Fails if Issuer claim does not match")
	validate.Flags().StringVar(&validateSub, "sub", "", "Fails if Subject claim does not match")
	validate.Flags().StringSliceVar(&validateAud, "aud", nil, "Fails if Audience claim does not match")
	validate.Flags().StringVar(&validateExp, "exp", "", "Fails if Expiration Time claim is before this value")
	validate.Flags().StringVar(&validateNbf, "nbf", "", "Fails of Not Before claim is after this value")
	validate.Flags().StringVar(&validateIat, "iat", "", "Fails if Issued At claim is after this value")
	validate.Flags().StringVar(&validateJti, "jti", "", "Fails if JWT ID claim does not match")
	validate.Flags().StringVarP(&validateAlgorithm, "algorithm", "a", "", "The algorithm to validate against. Fails on mismatch")
	validate.Flags().StringVarP(&validateSecret, "secret", "s", "", "The secret or public key (string, @file, or '-' to read from stdin). Fails if signature is invalid")
	root.AddCommand(validate)
}
